package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Объект для строительства
type House struct {
	size  string
	color string
}

// Интерфейс строителя объявляет все возможные этапы и шаги по строительству дома
type Builder interface {
	setSize(size string)
	setColor(color string)
	// build() House
}

// Все конкретные строители реализуют общий интерфейс по своему
type HouseBuilder struct {
	house House
}

func (b *HouseBuilder) setSize(size string) {
	b.house.size = size
}

func (b *HouseBuilder) setColor(color string) {
	b.house.color = color
}

// func (b *HouseBuilder) build() House {
// 	return House{
// 		size:  b.house.size,
// 		color: b.house.size,
// 	}
// }

// Директор знает в какой последовательности нужно заставлять работать строителя, чтобы получить ту или иную версию дома
type director struct {
	builder HouseBuilder
}

func (d *director) constructHouse(size string, color string) House {
	d.builder.setColor(color)
	d.builder.setSize(size)
	return d.builder.house
}

// func main() {
// 	// Создаем объект структуры на строителя дома
// 	builder := HouseBuilder{}
// 	// Cоздаем дом который будет строится через абстрактного строителя
// 	// house := builder.setSize("large").setColor("blue").build
// 	builder.setColor("size")
// 	builder.setSize("large")
// 	house := builder.build()
// 	fmt.Println(house)

// }
