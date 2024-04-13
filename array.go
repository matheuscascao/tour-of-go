package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "aa"
	a[1] = "bb"

	fmt.Println(a[0], a[1])
	fmt.Println(a)
}
