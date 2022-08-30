package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"

	"wildberries_traineeship/internal/models"
)

func setUuid(order models.OrderData) ([]byte, error) {
	order.OrderUid = uuid.New().String()

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
