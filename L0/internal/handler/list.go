package handler

import (
	"context"
	"net/http"
)

type OrderListHandler struct {
	Service OrderListService
}

type OrderListService interface {
	GetOrderList(ctx context.Context) (*[]string, error)
}

func (h *OrderListHandler) Method() string {
	return http.MethodGet
}

func (h *OrderListHandler) Path() string {
	return "/orders"
}

func (h *OrderListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	list, err := h.Service.GetOrderList(r.Context())
	if err != nil {
		writeResponse(w, r, err)
		return
	}
	writeResponse(w, r, list)
}
