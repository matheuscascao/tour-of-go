package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

type Book struct {
	Name string
}

func (b Book) String() string {
	return "Book test"
}

type Perimeter interface {
}

type Stringer interface {
	String() string
}

func main() {
	a := Article{
		Title:  "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}

	b := Book{
		Name: "test",
	}
	Print(a)
	Print(b)
}

func Print(s Stringer) {
	fmt.Println(s.String())
}
