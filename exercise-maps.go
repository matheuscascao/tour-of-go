package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	mapReturn := make(map[string]int)
	splited := strings.Split(s, " ")

	for _, v := range splited {
		mapReturn[v] = mapReturn[v] + 1
	}

	return mapReturn
}

func main() {
	testReturn := WordCount("I am learning Go!")
	fmt.Println(testReturn)
}
