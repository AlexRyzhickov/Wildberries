package handler

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"reflect"
	"strings"
	"wildberries_traineeship/internal/models"
	"wildberries_traineeship/internal/service"
)

type OrderHandler struct {
	Service OrderService
}

type OrderService interface {
	GetOrderInfo(ctx context.Context, ticker string) (*models.Order, error)
}

func (h *OrderHandler) Method() string {
	return http.MethodGet
}

func (h *OrderHandler) Path() string {
	return "/stock/{ticker}"
}

func (h *OrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ticker := chi.URLParam(r, "ticker")

	stock, err := h.Service.GetOrderInfo(r.Context(), ticker)

	if err != nil {
		if nsi := (service.NoOrderInfo{}); errors.As(err, &nsi) {
			writeResponse(w, r, badRequest{nsi.Error()})
			return
		}
		writeResponse(w, r, err)
		return
	}

	var resp interface{} = stock
	if fields := r.FormValue("fields"); fields != "" {
		// Using reflection to leave in the response only requested fields

		m := make(map[string]interface{})
		rv := reflect.ValueOf(*stock)
		for _, field := range strings.Split(fields, ",") {
			m[field] = rv.FieldByName(field).Interface()
		}
		resp = m
	}

	writeResponse(w, r, resp)
}
