package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/
// Структура продукта
type Product struct {
	name  string
	count int
}

// Метод по добавлению информации в продукт
func (p *Product) GetProduct(name string, count int) {
	p.name = name
	p.count = count
}

// Структура оплаты заказа в продукте
type Payments struct {
	sum    float32
	status bool
}

// Метод по проведению оплаты заказа
func (p *Payments) DoPayment(sum float32) {
	p.sum = sum
	p.status = true
	fmt.Println("Оплата заказа произведена")
}

// Инвойс - это документ, объединяющий функции товарной накладной, акта выполненных работ и счета на оплату
type Invoice struct {
	status bool
}

// Метод инвойса, который получает информацию об оплате
func (i *Invoice) SendInvoice(payments Payments) {
	// Если оплата произвенеа успешно
	if payments.status {
		fmt.Println("Инвойс успешно создан")
		i.status = true
	}
	fmt.Println("Инвойс не создан")
	i.status = false
}

// Cтруктура заказа
type Order struct {
}

// Паттерн фасад в данном случае предоставляет простой интерфейс интерфейс по проведению заказа и срывает все этапы по работе с самим заказом
// Проведение заказа
func (o *Order) PlaceOrder(name string, count int) {
	// Cоздаем продукт в заказе
	prod := Product{}
	prod.GetProduct(name, count)

	// Проводим оплату заказа
	payment := Payments{}
	payment.DoPayment(100)

	// Формируем документ инвойс
	invoice := Invoice{}
	invoice.SendInvoice(payment)

}
