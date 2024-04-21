package main

import "fmt"

type Square struct {
	Side float32
}

type Rectangle struct {
	SideA, SideB float32
}

func (s *Square) Perimeter() float32 {
	return (s.Side) * 4
}

func (r *Rectangle) Perimeter() float32 {
	return (r.SideA + r.SideB) * 2
}

type HasPerimeter interface {
	Perimeter() float32
}

func returnTriplePerimeter(geometryItem HasPerimeter) float32 {
	return geometryItem.Perimeter() * 10
}

func main() {
	sq := &Square{5}
	rec := &Rectangle{5, 5}

	fmt.Println(returnTriplePerimeter(sq))
	fmt.Println(returnTriplePerimeter(rec))
}
