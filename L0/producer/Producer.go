package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"strings"
	"time"
	"wildberries_traineeship/internal/models"
)

func genUuid() string {
	uuidWithHyphen := uuid.New()
	fmt.Println(uuidWithHyphen)
	return strings.Replace(uuidWithHyphen.String(), "-", "", -1)
}

func setUuid(order models.OrderData) ([]byte, error) {
	order.OrderUid = genUuid()

	bytes, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func main() {
	order := models.OrderData{}
	fileBytes, _ := os.ReadFile("./model.json")
	err := json.Unmarshal(fileBytes, &order)

	if err != nil {
		log.Fatal(err)
	}

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	for {
		bytes, err := setUuid(order)
		if err != nil {
			continue
		}
		nc.Publish("foo", bytes)
		time.Sleep(2 * time.Second)
	}
}
