package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice(a)

	a = append(a, 1)
	printSlice(a)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
