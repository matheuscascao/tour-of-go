package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64}

func main() {
	for _, v := range pow {
		fmt.Printf("v: %d", v)
	}
}
