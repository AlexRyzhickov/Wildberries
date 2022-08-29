package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"time"
	"wildberries_traineeship/internal/models"
)

func main() {

	order := models.Order{} // Read errors caught by unmarshal
	fileBytes, _ := os.ReadFile("./model.json")
	err := json.Unmarshal(fileBytes, &order)

	bytes, err := json.Marshal(order)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Println("Unmarshal error")
		return
	}

	fmt.Println(order)

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	for {
		nc.Publish("foo", bytes)
		time.Sleep(2 * time.Second)
	}
}

//package main
//
//import (
//	"github.com/nats-io/nats.go"
//	"sync"
//)
//
//func Block() {
//	w := sync.WaitGroup{}
//	w.Add(1)
//	w.Wait()
//}
//
//func main() {
//
//	//sc, _ := stan.Connect("prod", "sub-1")
//	//defer sc.Close()
//	//
//	//sc.Subscribe("bestellungen", func(m *stan.Msg) {
//	//	fmt.Printf("Got: %s\n", string(m.Data))
//	//})
//
//	//Block()
//
//	nc, _ := nats.Connect(nats.DefaultURL)
//	defer nc.Close()
//	//
//	////nc.Subscribe()
//	//
//	nc.Subscribe("request", func(m *nats.Msg) {
//		m.Respond([]byte("answer is 42"))
//	})
//	//
//	Block()
//}
