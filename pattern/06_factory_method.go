package pattern

import "log"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Именованный тип для техники
type technique string

// Константы выпускаемой техники
const (
	Computer   technique = "Computer"
	Notebook   technique = "Notebook"
	Smartphone technique = "Smartphone"
)

// Интерфейс Техники
type Technique interface {
	Use() string // Каждую технику можно использовать
}

// Фабрика по созданию новой техники
type Creator interface {
	CreateProduct(action technique) Technique // Фабричный метод, создающий новую технику
}

// Структура конкретной фабрики по изготовлению техники, которая будет реализовывать интерфейс Creator
type ConcreteCreator struct{}

// Конструктор для ConcreteCreator
func NewCreator() Creator {
	return &ConcreteCreator{}
}

// Структура компьютера
type CreateComputer struct {
	technique string
}

// Использование компьютера
func (p *CreateComputer) Use() string {
	return p.technique
}

// Объект ноутбука
type CreateNotebook struct {
	technique string
}

// Использование ноутбука
func (p *CreateNotebook) Use() string {
	return p.technique
}

// Объект смартфона
type CreateSmartphone struct {
	technique string
}

// Использование смартфона
func (p *CreateSmartphone) Use() string {
	return p.technique
}

// Реализация фабричного метода по выпуску новой техники
func (p *ConcreteCreator) CreateProduct(typeOfTechnique technique) Technique {
	var product Technique

	switch typeOfTechnique {
	case Computer:
		product = &CreateComputer{string(typeOfTechnique)}
	case Notebook:
		product = &CreateNotebook{string(typeOfTechnique)}
	case Smartphone:
		product = &CreateSmartphone{string(typeOfTechnique)}
	default:
		log.Fatalln("Данная фабрика не может создавать данную технику")
	}

	return product
}
