package pattern

import (
	"fmt"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

Данный шаблон используется для того, чтобы создавать объект который может менять свое поведение в зависимости от внутреннего состояния

Позволяет избегать многочисленных условий if

State – все состояния который может иметь данный объект
Context – текущее состояние объекта

cфетовор может гореть 3 цветами
*/

// Интерфейс светофора
type svetoforInterface interface {
	ToBur() // светофор может гореть
}

// Типы состояния
type RedColor struct{}
type GreenColor struct{}
type YellowCollor struct{}

// Реализация интерфейса состояниями
func (RedColor) ToBur() {
	fmt.Println("Светофор горит красным цветом")
}

func (GreenColor) ToBur() {
	fmt.Println("Светофор горит зеленым цветом")
}

func (YellowCollor) ToBur() {
	fmt.Println("Светофор горит зеленым цветом")
}

// Структура с полем состояния
type svetoforStruct struct {
	state svetoforInterface
}

// func main() {
// 	svetofor := svetoforStruct{RedColor{}}
// 	svetofor.state.ToBur()
// }
