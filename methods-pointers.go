package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) SumXandY() float64 {
	// return math.Sqrt(v.X*v.X + v.Y*v.Y)
	sum := v.X + v.Y
	return sum
}

func (v *Vertex) Scale(f float64) { //Methods with pointers modify the original instance of the object, creating a copy of it
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.SumXandY())
}
