package main

import "fmt"

func main() {
	// i, j := 42, 2701

	test := 44

	p := &test
	fmt.Println(*p)

	test = 99
	fmt.Println(*p)

	// *i = 11
	// fmt.Println(*p)
	// p = &j
	// *p = *p /37
	// fmt.Println(j)
}
