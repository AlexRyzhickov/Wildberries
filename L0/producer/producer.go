package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"

	"wildberries_traineeship/internal/models"
)

func setUuid(order models.OrderData) ([]byte, string, error) {
	order.OrderUid = uuid.New().String()

	bytes, err := json.Marshal(order)
	if err != nil {
		return nil, "", err
	}

	return bytes, order.OrderUid, nil
}

func main() {
	order := models.OrderData{}
	fileBytes, _ := os.ReadFile(os.Getenv("FILE_PATH"))
	err := json.Unmarshal(fileBytes, &order)

	if err != nil {
		log.Fatal(err)
	}

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	for {
		bytes, uid, err := setUuid(order)
		if err != nil {
			continue
		}
		nc.Publish(os.Getenv("SUBJ"), bytes)
		fmt.Println("Publish msg with uid", uid)
		time.Sleep(2 * time.Second)
	}
}
