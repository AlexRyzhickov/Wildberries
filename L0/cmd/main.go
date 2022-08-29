package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"
	"wildberries_traineeship/internal/config"
	"wildberries_traineeship/internal/handler"
	"wildberries_traineeship/internal/models"
	"wildberries_traineeship/internal/service"

	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func connectDB(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DBConn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Handler interface {
	Method() string
	Path() string
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func registerHandler(router chi.Router, handler Handler) {
	router.Method(handler.Method(), handler.Path(), handler)
}

func connectionsClosedForServer(server *http.Server) chan struct{} {
	connectionsClosed := make(chan struct{})
	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt)
		defer signal.Stop(shutdown)
		<-shutdown

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()
		log.Println("Closing connections")
		if err := server.Shutdown(ctx); err != nil {
			log.Println(err)
		}
		close(connectionsClosed)
	}()
	return connectionsClosed
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := connectDB(cfg)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("failed to connect nats streaming", err)
	}
	defer nc.Close()

	_, err = nc.Subscribe("foo", func(m *nats.Msg) {

		order := models.Order2{}
		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			log.Println("Unmarshal data from nats-streaming error", err)
		}
		data := models.Order{
			Id: rand.Intn(10000),
		}

		var inInterface map[string]interface{}
		inrec, _ := json.Marshal(order)
		err = json.Unmarshal(inrec, &inInterface)

		if err != nil {
			log.Println("Marshal data from nats-streaming error", err)
		}

		fmt.Println(inInterface)
		//for field, val := range inInterface {
		//fmt.Println("KV Pair: ", field, val)
		//}

		err = data.Order2.Set(inInterface)
		if err != nil {
			log.Println("Q", err)
			return
		}

		err = db.FirstOrCreate(&data).Error

		if err != nil {
			log.Println("Q2", err)
			return
		}

		//db.Updates(&u)

		//order := models.Order{}
		//
		//err := json.Unmarshal(m.Data, &order)
		//
		//if err != nil {
		//	log.Println("Unmarshal data from nats-streaming error")
		//} else {
		//	err = db.FirstOrCreate(&order).Error
		//	if err != nil {
		//		//return &pb.AddContactResponse{Msg: fmt.Sprintf("%s, %v", addError, err.Error())}, err
		//	}
		//}

		//err := db.FirstOrCreate
		//fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	if err != nil {
		log.Fatal("failed to subscribe", err)
	}

	service := service.NewService(db)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.AllowAll().Handler)

	router.Group(func(router chi.Router) {
		//router.Use(cacheMiddleware)
		registerHandler(router, &handler.OrderHandler{Service: service})
		//registerHandler(router, &handler.CountriesHandler{Service: service})
		//registerHandler(router, &handler.SectorsHandler{Service: service})
		//registerHandler(router, &handler.IndustriesHandler{Service: service})
		//registerHandler(router, &handler.SearchHandler{Service: service})
		//registerHandler(router, &handler.ExportHandler{Service: service})
		//registerHandler(router, &handler.ImportHandler{Service: service})
		//registerHandler(router, &handler.CountHandler{Service: service})
		//registerHandler(router, &handler.TableHandler{Service: service})
	})

	addr := fmt.Sprintf(":%d", cfg.Port)
	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	connectionsClosed := connectionsClosedForServer(&server)
	log.Println("Server is listening on " + addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Println(err)
	}
	<-connectionsClosed

	//cfg, err := config.New()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//lis, err := net.Listen("tcp", ":"+cfg.Port)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//s := grpc.NewServer()
	//
	//db, err := connectDB(cfg)
	//if err != nil {
	//	log.Fatal("failed to connect database", err)
	//}
	////if err = initializeDB(db); err != nil {
	////	log.Fatal("failed to init `contact` table", err)
	////}
	//
	//srv := service.NewAddressBookService(db)
	//
	//go func() {
	//	mux := runtime.NewServeMux()
	//	if err := pb.RegisterAddressBookServiceHandlerServer(context.Background(), mux, srv); err != nil {
	//		log.Fatal(err)
	//	}
	//	if err = http.ListenAndServe(":"+cfg.ProxyPort, mux); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	//
	//pb.RegisterAddressBookServiceServer(s, service.NewAddressBookService(db))
	//log.Printf("Server is listening on: %v", lis.Addr())
	//if err = s.Serve(lis); err != nil {
	//	log.Fatalf("Failed to serve: %v", err)
	//}
	//cfg, err := config.New()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//lis, err := net.Listen("tcp", ":"+cfg.Port)
	//if err != nil {
	//	log.Fatal(err)
	//}
	////s := grpc.NewServer()
	//
	//db, err := connectDB(cfg)
	//if err != nil {
	//	log.Fatal("failed to connect database", err)
	//}
	//
	//nc, _ := nats.Connect(nats.DefaultURL)
	//defer nc.Close()
	//
	//nc.Subscribe("foo", func(m *nats.Msg) {
	//	fmt.Printf("Received a message: %s\n", string(m.Data))
	//})
	//
	//w := sync.WaitGroup{}
	//w.Add(1)
	//w.Wait()
}

//package main
//
//import (
//	"github.com/nats-io/nats.go"
//	"time"
//)
//
//func main() {
//
//	//sc, _ := stan.Connect("prod", "simple-pub")
//	//
//	//for i := 1; ; i++ {
//	//	sc.Publish("bestellungen", []byte("Bestellung "+strconv.Itoa(i)))
//	//	time.Sleep(2 * time.Second)
//	//}
//	//
//	//sc.Close()
//	nc, _ := nats.Connect(nats.DefaultURL)
//	defer nc.Close()
//
//	for i := 0; i < 100; i++ {
//		nc.Publish("foo", []byte("Hello World"))
//		time.Sleep(2 * time.Second)
//	}
//
//	// Simple Publisher
//
//	//// Simple Async Subscriber
//	//nc.Subscribe("foo", func(m *nats.Msg) {
//	//	fmt.Printf("Received a message: %s\n", string(m.Data))
//	//})
//	//
//	//// Responding to a request message
//	//nc.Subscribe("request", func(m *nats.Msg) {
//	//	m.Respond([]byte("answer is 42"))
//	//})
//	//
//	//// Simple Sync Subscriber
//	//sub, err := nc.SubscribeSync("foo")
//	//m, err := sub.NextMsg(timeout)
//	//
//	//// Channel Subscriber
//	//ch := make(chan *nats.Msg, 64)
//	//sub, err := nc.ChanSubscribe("foo", ch)
//	//msg := <-ch
//	//
//	//// Unsubscribe
//	//sub.Unsubscribe()
//	//
//	//// Drain
//	//sub.Drain()
//	//
//	//// Requests
//	//msg, err := nc.Request("help", []byte("help me"), 10*time.Millisecond)
//	//
//	//// Replies
//	//nc.Subscribe("help", func(m *nats.Msg) {
//	//	nc.Publish(m.Reply, []byte("I can help!"))
//	//})
//	//
//	//// Drain connection (Preferred for responders)
//	//// Close() not needed if this is called.
//	//nc.Drain()
//	//
//	//// Close connection
//	//nc.Close()
//
//}
