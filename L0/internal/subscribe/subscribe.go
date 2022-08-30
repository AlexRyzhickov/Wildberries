package subscribe

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
	"wildberries_traineeship/internal/models"
)

func ProcessOrder(db *gorm.DB, m *nats.Msg) error {
	orderData := models.OrderData{}
	err := json.Unmarshal(m.Data, &orderData)
	if err != nil {
		return err
	}
	data := models.Order{
		Id: orderData.OrderUid,
	}
	var inInterface map[string]interface{}
	inrec, err := json.Marshal(orderData)
	if err != nil {
		return err
	}
	err = json.Unmarshal(inrec, &inInterface)
	if err != nil {
		return err
	}
	err = data.OrderData.Set(inInterface)
	if err != nil {
		return err
	}
	err = db.FirstOrCreate(&data).Error
	if err != nil {
		return err
	}
	return nil
}
