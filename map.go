package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// var m map[string]Vertex

var testMap = map[string]int{
	"um":   1,
	"dois": 2,
}

func main() {
	// m = make(map[string]Vertex)
	// m["Bell"] = Vertex{1, 2}
	// fmt.Println(m["Bell"])

	testMap["tres"] = 3
	fmt.Println(testMap)

	v, ok := testMap["tres"]
	fmt.Println("The value", v, ok)
}
