package main

import "fmt"

type FoodDelivery interface {
	Deliver() string
}

type FoodDeliveryService struct {
	delivery FoodDelivery
}

func (a *FoodDeliveryService) Deliver() string {
	return a.delivery.Deliver()
}

func (a *FoodDeliveryService) SetDeliveryState(delivery FoodDelivery) {
	a.delivery = delivery
}

func NewFoodDelivery() *FoodDeliveryService {
	return &FoodDeliveryService{delivery: &DeliveryClub{}}
}

type DeliveryClub struct {
}

func (a *DeliveryClub) Deliver() string {
	return "DeliveryClub: Быстро всё доставим ! Всё будет очень вкусно !"
}

type YandexEda struct{}

func (a *YandexEda) Deliver() string {
	return "Яндекс.Еда: Ой, ой. У нас очень высокая загруженность курьеров, и ваш заказ кажется съел наш курьер, вот вам промокодик, не расттраивайтесь ;)"
}

func main() {
	y := &YandexEda{}
	fd := NewFoodDelivery()
	fmt.Println(fd.Deliver())
	fd.SetDeliveryState(y)
	fmt.Println(fd.Deliver())
}

/*
Состояние — это поведенческий паттерн, позволяющий динамически изменять поведение объекта при смене его состояния.
*/
