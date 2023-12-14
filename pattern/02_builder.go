package main

import "fmt"

/*
Строитель - паттерн, порождающий объекты.
Он отделяет конструирование сложного объекта от его представления,
так что в результате одного и того же процесса конструирования могут получаться разные представления.

Применимость паттерна:
- алгоритм создания сложного объекта не должен зависеть от того,
из каких частей состоит объект и как они стыкуются между собой;
- процесс конструирования должен обеспечивать различные представления конструируемого объекта.

Плюсы и минусы:
- позволяет изменять внутреннее представление продукта;
- изолирует код, реализующий конструирование и представление;
- дает более тонкий контроль над процессом конструирования.

Примеры использования:
- приложение для конвертирования из формата RTF из библиотеки ЕТ++ [WGM88]
*/

// Объект, который будет конструироваться
type house struct {
	windowType string
	doorType   string
	floor      int
}

// Интерфейс строитель с методами шагов конструирования объекта
// и методом, возвращающим объект
type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(builderType string) iBuilder {
	switch builderType {
	case "normal":
		return &normalBuilder{}
	case "igloo":
		return &iglooBuilder{}
	default:
		return nil
	}
}

// Директор, управляющий конструированием
type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{b}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

// Строитель 1
type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *normalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *normalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *normalBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// Строитель 2
type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *iglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *iglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *iglooBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// Пример
func main() {
	fmt.Println("Builder example")

	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Println(normalHouse)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Println(iglooHouse)
}
