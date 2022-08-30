package handler

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
	"wildberries_traineeship/internal/models"
)

type OrderHandler struct {
	Service OrderService
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

	order, err := h.Service.GetOrderInfo(r.Context(), id)

	if err != nil {
		writeResponse(w, r, err)
		return
	}

	writeResponse(w, r, order)
}
