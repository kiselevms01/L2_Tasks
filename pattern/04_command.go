package main

import "fmt"

/*
Команда - паттерн поведения объектов.
Он инкапсулирует запрос как объект, позволяя тем самым задавать параметры клиентов
для обработки соответствующих запросов, ставить запросы в очередь или протоколировать их,
а также поддерживать отмену операций.

Применимость паттерна:
- когда необходимо параметризовать объекты выполняемым действием;
- когда необходимо определять, ставить в очередь и выполнять запросы в разное время;
- когда необходимо поддержать отмену операций;
- когда необходимо поддержать протоколирование изменений,
чтобы их можно было выполнить повторно после аварийной остановки системы;
- когда необходимо структурировать систему на основе высокоуровневых операций, построенных из примитивных

Плюсы и минусы:
- позволяет добиться высокой гибкости при проектировании пользовательского интерфейса.

Примеры использования:
- в системе МасАрр [Арр89] команды широко применяются для реализации допускающих отмену операций.
*/

// Отправитель содержит команду
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

// Интерфейс команды
type command interface {
	execute()
}

// Команда 1
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

// Команда 2
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

// Интерфейс получателя
type device interface {
	on()
	off()
}

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
}

func (t *tv) off() {
	t.isRunning = false
}

func main() {
	fmt.Println("Command example")

	// Полчатель
	tv := &tv{}

	// Команды
	onCommand := &onCommand{
		device: tv,
	}

	offCommand := &offCommand{
		device: tv,
	}

	// Отправители с разными командами
	onButton := &button{
		command: onCommand,
	}
	onButton.press()
	fmt.Println(tv.isRunning)

	offButton := &button{
		command: offCommand,
	}
	offButton.press()
	fmt.Println(tv.isRunning)
}
