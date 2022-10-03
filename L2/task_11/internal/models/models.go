package models

const (
	Day = iota
	Week
	Month
)

type Event struct {
	UserId  int    `json:"user_id"`
	EventId int    `json:"event_id"`
	Date    string `json:"date"`
	Event   string `json:"event"`
}

type Result struct {
	Result interface{} `json:"result"`
}
