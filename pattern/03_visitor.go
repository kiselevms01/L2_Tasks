package main

import (
	"fmt"
	"math"
)

/*
Посетитель - паттерн поведения объектов.
Он описывает операцию, выполняемую с каждым объектом из некоторой структуры.
Паттерн посетитель позволяет определить новую операцию, не изменяя классы этих объектов.

Применимость паттерна:
- в структуре присутствуют объекты многих классов с различными интерфейсами
и есть необходимость выполнять над ними операции, зависящие от конкретных классов;
- над объектами, входящими в состав структуры, надо выполнять разнообразные,
не связанные между собой операции и вы не хотите «засорять» классы такими операциями;
- классы, устанавливающие структуру объектов, изменяются редко,
но новые операции над этой структурой добавляются часто.

Плюсы и минусы:
- упрощает добавление новых операций;
- объединяет родственные операции и отсекает те, которые не имеют к ним отношения;
- добавление новых классов ConcreteElement затруднено;
- посещение различных иерархий классов;
- аккумулирование состояния;
- нарушение инкапсуляции.

Примеры использования:
- В компиляторе Smalltalk-80 имеется класс посетителя, который называется ProgramNodeEnumerator.
В основном он применяется в алгоритмах анализа исходного текста программы
и не используется ни для генерации кода, ни для красивой печати, хотя мог бы.
*/

type shape interface {
	accept(visitor)
}

// Тип объекта 1
type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

// Тип объекта 2
type circle struct {
	radius int
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

// Тип объекта 3
type rectangle struct {
	l int
	b int
}

func (t *rectangle) accept(v visitor) {
	v.visitForrectangle(t)
}

// Интерфейс visitor с методами для каждого типа объекта
type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForrectangle(*rectangle)
}

// Имплементация интерфейса visitor
type areaCalculator struct {
	area float64
}

func (a *areaCalculator) visitForSquare(s *square) {
	a.area = float64(s.side * s.side)
}

func (a *areaCalculator) visitForCircle(s *circle) {
	a.area = float64(math.Pi) * float64(s.radius*s.radius)
}
func (a *areaCalculator) visitForrectangle(s *rectangle) {
	a.area = float64(s.b * s.l)
}

func main() {
	fmt.Println("Visitor example")
	fmt.Println()

	square := &square{side: 2}
	circle := &circle{radius: 3}
	rectangle := &rectangle{l: 2, b: 3}

	areaCalculator := &areaCalculator{}

	square.accept(areaCalculator)
	fmt.Printf("square area is %f\n", areaCalculator.area)
	circle.accept(areaCalculator)
	fmt.Printf("circle area is %f\n", areaCalculator.area)
	rectangle.accept(areaCalculator)
	fmt.Printf("rectangle area is %f\n", areaCalculator.area)
}
