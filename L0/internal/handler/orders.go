package handler

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
	"wildberries_traineeship/internal/cache"
	"wildberries_traineeship/internal/models"
)

const recordNotFoundMsg = "record not found"

type OrderHandler struct {
	Service OrderService
	Cache   cache.Cache
}

type OrderService interface {
	GetOrderInfo(ctx context.Context, ticker string) (*models.OrderData, error)
}

func (h *OrderHandler) Method() string {
	return http.MethodGet
}

func (h *OrderHandler) Path() string {
	return "/order/{id}"
}

func (h *OrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if item, isHas := h.Cache.Get(id); isHas {
		writeResponse(w, r, item)
		return
	}
	order, err := h.Service.GetOrderInfo(r.Context(), id)
	if err != nil {
		if err.Error() == recordNotFoundMsg {
			writeResponse(w, r, fmt.Sprintf("Record with id %s not found", id))
			return
		}
		writeResponse(w, r, err)
		return
	}
	h.Cache.Set(id, order, 30*time.Minute)
	writeResponse(w, r, order)
}
