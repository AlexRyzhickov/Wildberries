package utils

import "wildberries_traineeship/internal/models"

func ExtractOrderData(order models.Order) (*models.OrderData, error) {
	orderData := models.OrderData{}
	err := order.OrderData.AssignTo(&orderData)
	if err != nil {
		return nil, err
	}
	return &orderData, nil
}
