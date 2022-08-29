package models

import "github.com/jackc/pgtype"

type Order struct {
	Id     int
	Order2 pgtype.JSONB `gorm:"column:info"`
}

//
//type JSONB map[string]interface{}
//
//func (j JSONB) Value() (driver.Value, error) {
//	valueString, err := json.Marshal(j)
//	return string(valueString), err
//}
//
//func (j *JSONB) Scan(value interface{}) error {
//	if err := json.Unmarshal(value.([]byte), &j); err != nil {
//		return err
//	}
//	return nil
//}

type Order2 struct {
	OrderUid          string   `json:"order_uid,omitempty"`
	TrackNumber       string   `json:"track_number,omitempty"`
	Entry             string   `json:"entry,omitempty"`
	Delivery          Delivery `json:"delivery,omitempty"`
	Payment           Payment  `json:"payment,omitempty"`
	Items             []Item   `json:"items,omitempty"`
	Locale            string   `json:"locale,omitempty"`
	InternalSignature string   `json:"internal_signature,omitempty"`
	CustomerId        string   `json:"customer_id,omitempty"`
	DeliveryService   string   `json:"delivery_service,omitempty"`
	ShardKey          string   `json:"shardkey,omitempty"`
	SmId              int      `json:"sm_id,omitempty"`
	DateCreated       string   `json:"date_created,omitempty"`
	OofShard          string   `json:"oof_shard,omitempty"`
}

type Delivery struct {
	Name    string `json:"name,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Zip     string `json:"zip,omitempty"`
	City    string `json:"city,omitempty"`
	Address string `json:"address,omitempty"`
	Region  string `json:"region,omitempty"`
	Email   string `json:"email,omitempty"`
}

type Payment struct {
	Transaction  string `json:"transaction,omitempty"`
	RequestId    string `json:"request_id ,omitempty"`
	Currency     string `json:"currency,omitempty"`
	Provider     string `json:"provider,omitempty"`
	Amount       int    `json:"amount,omitempty"`
	PaymentDt    int    `json:"payment_dt,omitempty"`
	Bank         string `json:"bank,omitempty"`
	DeliveryCost int    `json:"delivery_cost,omitempty"`
	GoodsTotal   int    `json:"goods_total,omitempty"`
	CustomFee    int    `json:"custom_fee,omitempty"`
}

type Item struct {
	ChrtId      int    `json:"chrt_id,omitempty"`
	TrackNumber string `json:"track_number,omitempty"`
	Price       int    `json:"price,omitempty"`
	Rid         string `json:"rid,omitempty"`
	Name        string `json:"name,omitempty"`
	Sale        int    `json:"sale,omitempty"`
	Size        string `json:"size,omitempty"`
	TotalPrice  int    `json:"total_price,omitempty"`
	NmId        int    `json:"nm_id,omitempty"`
	Brand       string `json:"brand,omitempty"`
	Status      int    `json:"status,omitempty"`
}
