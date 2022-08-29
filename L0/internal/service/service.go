package service

import (
	"context"
	"fmt"
	_ "github.com/nats-io/nats.go"
	"gorm.io/gorm"
	"wildberries_traineeship/internal/models"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s Service) GetOrderInfo(ctx context.Context, ticker string) (*models.Order, error) {

	return &models.Order{Id: 123}, nil
}

type NoOrderInfo struct {
	ticker string
}

func (err NoOrderInfo) Error() string {
	return fmt.Sprintf("No info about the stock with ticker %v", err.ticker)
}
