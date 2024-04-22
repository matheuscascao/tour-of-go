package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Int hehe", v)
	case string:
		fmt.Printf("String hehe", v)
	default:
		fmt.Printf("idk")
	}
}

func main() {
	do(21)
	do("hello")
	do(313.12)
}
