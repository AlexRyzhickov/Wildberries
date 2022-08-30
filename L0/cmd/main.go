package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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

		order := models.OrderData{}
		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			log.Println("Unmarshal data from nats-streaming error", err)
		}
		data := models.Order{
			Id: order.OrderUid,
		}

		var inInterface map[string]interface{}
		inrec, _ := json.Marshal(order)
		err = json.Unmarshal(inrec, &inInterface)

		if err != nil {
			log.Println("Marshal data from nats-streaming error", err)
		}

		fmt.Println(inInterface)

		err = data.OrderData.Set(inInterface)
		if err != nil {
			log.Println("Q", err)
			return
		}

		err = db.FirstOrCreate(&data).Error

		if err != nil {
			log.Println("Q2", err)
			return
		}

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
	})

	addr := fmt.Sprintf(":%d", cfg.Port)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	connectionsClosed := connectionsClosedForServer(&server)
	log.Println("Server is listening on " + addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Println(err)
	}
	<-connectionsClosed
}
