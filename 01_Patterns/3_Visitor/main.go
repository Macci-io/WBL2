package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Краткое описание:
	Паттерн "посетитель" позволяет нам добавить функционала к существующей структуре, не изменяя ее структуру.
Плюсы:
	1. Упрощает добавление операций, работающих со сложными структурами объектов.
	2. Объединяет родственные операции в одном классе.
	3. Посетитель может накапливать состояние при обходе структуры элементов.
Минусы:
	1. Паттерн не оправдан, если иерархия элементов часто меняется.
	2. Может привести к нарушению инкапсуляции элементов.
*/

func main() {
	shapes := NewShapes()
	visitorShape := VisitorShapeToPrint{}
	shapes.Visit(visitorShape)
}

// IVisitorShape Интерфейс для будущих посетителей
type IVisitorShape interface {
	VisitRectangle(rectangle *Rectangle)
	VisitCircle(circle *Circle)
}

// VisitorShapeToPrint Структура реализующая интерфейс выше
type VisitorShapeToPrint struct{}

//VisitRectangle посещает объект типа Rectangle
func (VisitorShapeToPrint) VisitRectangle(rectangle *Rectangle) {
	fmt.Printf("Visitor visited Rectangle: %+v\n", *rectangle)
}

//VisitCircle посещает объект типа Circle
func (VisitorShapeToPrint) VisitCircle(circle *Circle) {
	fmt.Printf("Visitor visited Circle: %+v\n", *circle)
}

// Shapes наша структура с возможностью приема посетителей
type Shapes struct {
	rectangle Rectangle
	circle    Circle
}

// Rectangle объект прямоугольника
type Rectangle struct{ width, height float32 }

// Circle объект круга
type Circle struct{ radius float32 }

// Visit 	---> Принимает интерфейс посетителя <---
func (s *Shapes) Visit(v IVisitorShape) {
	v.VisitRectangle(&s.rectangle)
	v.VisitCircle(&s.circle)
}

// NewShapes простой генератор случайных фигур
func NewShapes() *Shapes {
	s := new(Shapes)
	rand.Seed(time.Now().UnixNano())
	s.rectangle.width = rand.Float32() * 10
	s.rectangle.height = rand.Float32() * 10
	s.circle.radius = rand.Float32() * 10
	return s
}
