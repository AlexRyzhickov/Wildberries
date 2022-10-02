package service

import (
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
	"wildberries_traineeship/internal/models"
)

const layout = "2006-01-02"

type Service struct {
	m sync.Map
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateEvent(event models.Event) {
	s.m.Store(key(event.UserId, event.EventId), event)
}

func (s *Service) UpdateEvent(event models.Event) {
	s.m.Store(key(event.UserId, event.EventId), event)
}

func (s *Service) DeleteEvent(event models.Event) {
	s.m.Delete(key(event.UserId, event.EventId))
}

func (s *Service) GetEvents(userId int, period int) ([]models.Event, error) {
	strUserId := strconv.Itoa(userId)
	events := make([]models.Event, 0)
	today := time.Now()
	s.m.Range(func(key, value interface{}) bool {
		eventUserId := strings.Split(key.(string), "&")[0]
		if eventUserId == strUserId {
			event := value.(models.Event)
			date, err := time.Parse(layout, event.Date)
			if err == nil {
				if affiliation(period, today, date) {
					events = append(events, event)
				}
			} else {
				log.Println(err)
			}
		}
		return true
	})
	return events, nil
}

func affiliation(period int, today, date time.Time) bool {
	switch period {
	case models.Day:
		return today.Year() == date.Year() && today.Month() == date.Month() && today.Day() == date.Day()
	case models.Week:
		if today.Year() == date.Year() {
			_, todayWeek := today.ISOWeek()
			_, dateWeek := date.ISOWeek()
			if todayWeek == dateWeek {
				return true
			}
		}
		return false
	case models.Month:
		return today.Year() == date.Year() && today.Month() == date.Month()
	}
	return false
}

func key(userId, eventId int) string {
	return strconv.Itoa(userId) + "&" + strconv.Itoa(eventId)
}
