package main

import "fmt"

type Sword struct {
	strike StrikeOn
}

func (a *Sword) CauseDamage() {
	a.strike.takeDamage(100)
}

type StrikeOn interface {
	takeDamage(damage int)
}

type Swordsman struct {
	personage Personage
	lives     int
}

func (s *Swordsman) takeDamage(damage int) {
	if damage < s.lives {
		s.lives = s.lives - damage
		s.personage.live()
	} else {
		s.personage.die()
	}
}

type Personage interface {
	die()
	live()
}

type ImmortalWizard struct {
	personage Personage
}

func (c *ImmortalWizard) takeDamage(damage int) {
	c.personage.live()
}

type Elf struct{}

func (e *Elf) die() {
	fmt.Println("Elf died for our Lady Galadriel !")
}

func (e *Elf) live() {
	fmt.Println("I'm Elf, i'm fighting for the light !")
}

type Dwarf struct{}

func (d *Dwarf) die() {
	fmt.Println("I'm Dwarf, I die with honor !")
}

func (d *Dwarf) live() {
	fmt.Println("Dwarves never give up !")
}

func main() {
	dwarf := Dwarf{}
	elf := Elf{}
	immortalWizard := ImmortalWizard{personage: &elf}
	swordsman := Swordsman{personage: &dwarf, lives: 50}
	sword1 := Sword{strike: &immortalWizard}
	sword2 := Sword{strike: &swordsman}
	sword1.CauseDamage()
	sword2.CauseDamage()
}

/*
Команда — это поведенческий паттерн, позволяющий заворачивать запросы или простые операции в отдельные объекты.
+ Паттерн превращает запросы в объекты, позволяя передавать их как аргументы при вызове методов.
+ Паттерн позволяет хранить историю операций команд, производить операцию отката или отмены.
*/
