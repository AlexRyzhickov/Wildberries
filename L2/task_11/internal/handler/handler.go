package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"wildberries_traineeship/internal/models"
)

type CreateEventService interface {
	CreateEvent(event models.Event)
}

type CreateEventHandler struct {
	Service CreateEventService
}

func (h CreateEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, Error{err.Error()})
		return
	}
	h.Service.CreateEvent(event)
	writeResponse(w, http.StatusOK, models.Result{Result: event})
}

type UpdateEventService interface {
	UpdateEvent(event models.Event)
}

type UpdateEventHandler struct {
	Service UpdateEventService
}

func (h UpdateEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, Error{err.Error()})
		return
	}
	h.Service.UpdateEvent(event)
	writeResponse(w, http.StatusOK, models.Result{Result: event})
}

type DeleteEventService interface {
	DeleteEvent(event models.Event)
}

type DeleteEventHandler struct {
	Service DeleteEventService
}

func (h DeleteEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, Error{err.Error()})
		return
	}
	h.Service.DeleteEvent(event)
	writeResponse(w, http.StatusOK, models.Result{Result: event})
}

type EventsService interface {
	GetEvents(userId int, period int) ([]models.Event, error)
}

type EventsHandler struct {
	Service EventsService
	Period  int
}

func (h EventsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Println(err)
		writeResponse(w, http.StatusBadRequest, Error{err.Error()})
		return
	}
	events, err := h.Service.GetEvents(userId, h.Period)
	if err != nil {
		log.Println(err)
		writeResponse(w, http.StatusInternalServerError, Error{err.Error()})
		return
	}
	writeResponse(w, http.StatusInternalServerError, models.Result{Result: events})
}
