package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type StringerTest interface {
	String() string
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func PrintString(p StringerTest) string {
	fmt.Printf(p.String())
	return "a"
}

func main() {
	a := Person{"Arthur Dent", 42}
	// fmt.Println(a.String())

	fmt.Println("")

	PrintString(a)

}
