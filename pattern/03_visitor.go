package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

	// Расширение функционала структуры без изменения самой структуры
*/

type Car interface {
	accept(visitor)
}

type Tesla struct{}
type Bmw struct{}
type Audi struct{}

func (Tesla) accept(v visitor) {
	v.serviceTesla()

}

func (Bmw) accept(v visitor) {
	v.serviceBmw()
}

func (Audi) accept(v visitor) {
	v.serviceAuid()
}

type visitor interface {
	serviceTesla()
	serviceBmw()
	serviceAuid()
}

type mechanic struct{}

func (mechanic) serviceTesla() {
	fmt.Println("Обслуживание Tesla")
}

func (mechanic) serviceBmw() {
	fmt.Println("Обслуживание BMW")
}

func (mechanic) serviceAudi() {
	fmt.Println("Обслуживание Audi")
}
