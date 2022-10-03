package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wildberries_traineeship/internal/handler"
	"wildberries_traineeship/internal/models"
	"wildberries_traineeship/internal/service"
)

func main() {
	srv := http.Server{
		Addr: ":8080",
	}

	s := service.NewService()
	http.Handle("/create_event", loggingMiddleware(handler.CreateEventHandler{Service: s}.ServeHTTP))
	http.Handle("/update_event", loggingMiddleware(handler.UpdateEventHandler{Service: s}.ServeHTTP))
	http.Handle("/delete_event", loggingMiddleware(handler.DeleteEventHandler{Service: s}.ServeHTTP))
	http.Handle("/events_for_day", loggingMiddleware(handler.EventsHandler{Service: s, Period: models.Day}.ServeHTTP))
	http.Handle("/events_for_week", loggingMiddleware(handler.EventsHandler{Service: s, Period: models.Week}.ServeHTTP))
	http.Handle("/events_for_month", loggingMiddleware(handler.EventsHandler{Service: s, Period: models.Month}.ServeHTTP))

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(shutdown)

	go func() {
		log.Println("Server is listening on :" + srv.Addr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-shutdown

	log.Println("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Server stopped gracefully")
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, r.RemoteAddr, r.UserAgent())
		next(w, r)
	}
}
