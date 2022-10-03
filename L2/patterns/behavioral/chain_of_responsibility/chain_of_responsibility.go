package main

import "fmt"

type Border interface {
	pass(*RussianProgrammer)
	setNextBorder(Border)
}

type RussianBorder struct {
	next Border
}

func (b *RussianBorder) pass(p *RussianProgrammer) {
	if p.russianBorderDone {
		fmt.Println("Russian Programmer already passed russian border")
		b.next.pass(p)
		return
	}
	fmt.Println("Russian Programmer passed russian border")
	p.russianBorderDone = true
	b.next.pass(p)
}

func (b *RussianBorder) setNextBorder(next Border) {
	b.next = next
}

type GeorgianBorder struct {
	next Border
}

func (b *GeorgianBorder) pass(p *RussianProgrammer) {
	if p.georgianBorderDone {
		fmt.Println("Russian Programmer already passed georgian border")
		b.next.pass(p)
		return
	}
	fmt.Println("Russian Programmer passed georgian border")
	p.georgianBorderDone = true
	b.next.pass(p)
}

func (b *GeorgianBorder) setNextBorder(next Border) {
	b.next = next
}

type SerbianBorder struct {
	next Border
}

func (b *SerbianBorder) pass(p *RussianProgrammer) {
	if p.SerbianBorderDone {
		fmt.Println("Russian Programmer already passed serbian border")
		b.next.pass(p)
		return
	}
	fmt.Println("Russian Programmer passed serbian border")
	p.SerbianBorderDone = true
	fmt.Println("Welcome to Serbia !")
}

func (c *SerbianBorder) setNextBorder(next Border) {
	c.next = next
}

type RussianProgrammer struct {
	russianBorderDone  bool
	georgianBorderDone bool
	SerbianBorderDone  bool
}

func main() {
	serbianBorder := &SerbianBorder{}

	georgianBorder := &GeorgianBorder{}
	georgianBorder.setNextBorder(serbianBorder)

	russianBorder := &RussianBorder{}
	russianBorder.setNextBorder(georgianBorder)

	russianProgrammer := &RussianProgrammer{}
	russianBorder.pass(russianProgrammer)
}

/*
Цепочка обязанностей — это поведенческий паттерн, позволяющий передавать запрос по цепочке потенциальных обработчиков, пока один из них не обработает запрос.
+ Цепочка обязанностей позволяет выстраивать цепочки обработчиков объектов.
*/
