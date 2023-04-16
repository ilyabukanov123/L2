package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Объект для строительства
type House struct {
	Size  string
	Color string
}

// Интерфейс строителя объявляет все возможные этапы и шаги по строительству дома
type Builder interface {
	setSize(size string) HouseBuilder
	setColor(color string) HouseBuilder
	build() House
}

// Все конкретные строители реализуют общий интерфейс по своему
type HouseBuilder struct {
	size  string
	color string
}

func (b HouseBuilder) setSize(size string) HouseBuilder {
	b.size = size
	return b
}

func (b HouseBuilder) setColor(color string) HouseBuilder {
	b.color = color
	return b
}

func (b HouseBuilder) build() House {
	return House{
		Size:  b.size,
		Color: b.color,
	}
}

// Директор знает в какой последовательности нужно заставлять работать строителя, чтобы получить ту или иную версию дома
type director struct {
	builder HouseBuilder
}

func (d *director) constructHouse(size string, color string) House {
	d.builder = d.builder.setColor(color)
	d.builder = d.builder.setSize(size)
	return d.builder.build()
}

// Реализация без директора
// func main() {
// 	// Создаем объект структуры на строителя дома
// 	builder := HouseBuilder{}
// 	house := builder.setColor("red").setSize("large").build

// 	fmt.Println(house)

// }

// Реализация с директором
// func main() {
// 	builder := HouseBuilder{}
// 	director := director{builder}
// 	fmt.Println(director.constructHouse("large", "red"))
// }
