package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

	Данный паттерн позволяет передавать выполнение последовательно по цепочки

	Плюс: уменьшает зависимость между клиентами и обработчиком, принцип единственной обязанности, принцип открытости и закрытости
	Минусы: запрос может отсаться не обработанным
*/

// Общий интерфейс для логирования
type logger interface {
	print(message) // Метод по выводу сообщения об ошибке
}

// Сообщение об ошибке логирования
type message struct {
	level   uint   // Уроверь логирования
	payload string //  Сообщение об ошибке
}

// Уровни логирования
type infoLogger struct{ next logger }
type warningLogger struct{ next logger }
type errorLogger struct{ next logger }
type panicLogger struct{}

// Методы вывода информативной ошибки
func (l *infoLogger) print(msg message) {
	if msg.level == 0 {
		fmt.Println("INFO: ", msg.payload)
		return
	}
	l.next.print(msg)
}

// Метод вывода предупреждения
func (l *warningLogger) print(msg message) {
	if msg.level == 1 {
		fmt.Println("WARN: ", msg.payload)
		return
	}
	l.next.print(msg)
}

// Метод вывода ошибки
func (l *errorLogger) print(msg message) {
	if msg.level == 2 {
		fmt.Println("ERR: ", msg.payload)
		return
	}
	l.next.print(msg)
}

// Метод вывода критической ошибки
func (l *panicLogger) print(msg message) {
	fmt.Println("PANIC: ", msg.payload)
}

/*
	func main() {
	msg := message{3, "log message"}
	panic := panicLogger{}
	err := errorLogger{&panic}
	warn := warningLogger{&err}
	info := infoLogger{&warn}
	info.print(msg)
	}
*/
