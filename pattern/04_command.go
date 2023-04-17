package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// Коммутатор
type switcher struct {
	isOn bool // активен или не активен
}

// Метод включение коммутатора
func (s *switcher) on() {
	s.isOn = true
}

// Метод выключение коммутатора
func (s *switcher) off() {
	s.isOn = false
}

// Кнопка по реализации команды
type button struct {
	cmd command
}

// Метод по установлению команды к кнопке
func (b *button) setCommand(cmd command) {
	b.cmd = cmd
}

// Метод по работе с коммутатором(выключение или включение)
func (b *button) executeCommang() {
	b.cmd.execute()
}

// интерфейс команды
type command interface {
	execute()
}

// Команда которая будет работать с коммутатором и включать его
type switcherOnCommand struct{ s *switcher }

// Команды на включение коммутатора
func (c *switcherOnCommand) execute() {
	c.s.on()
}

// Команда которая будет работать с коммутатором и выключать его
type switcherOffCommand struct {
	s *switcher
}

// Команда на выключение коммутатора
func (c *switcherOffCommand) execute() {
	c.s.off()
}

// func main() {
// 	button := button{}
// 	switcher := switcher{}

// 	onCommang := switcherOnCommand{&switcher}
// 	button.setCommand(&onCommang)
// 	button.executeCommang()
// 	fmt.Println(switcher)

// 	offCommang := switcherOffCommand{&switcher}
// 	button.setCommand(&offCommang)
// 	button.executeCommang()
// 	fmt.Println(switcher)
// }
