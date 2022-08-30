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

func (s Service) GetOrderInfo(ctx context.Context, ticker string) (*models.OrderData, error) {
	order := models.Order{
		Id: "b563feb7b2b84b6test",
	}

	//result := map[string]interface{}{}
	s.db.Model(&models.Order{}).First(&order)

	order2 := models.OrderData{}

	err := order.OrderData.AssignTo(&order2)
	if err != nil {
		fmt.Println("!!!!!!!!!!!!!!!")
		//t.Fatal(err)
	}

	//result := map[string]interface{}{}
	//s.db.Table("orders").First(&result)

	fmt.Println("hi", order.Id, order2)

	//err := s.db.Find(order).Error
	//
	//if err != nil {
	//	return nil, err
	//}

	return &order2, nil
}

type NoOrderInfo struct {
	ticker string
}

func (err NoOrderInfo) Error() string {
	return fmt.Sprintf("No info about the stock with ticker %v", err.ticker)
}
