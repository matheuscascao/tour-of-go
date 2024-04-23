package main

import (
	"fmt"
)

type ErrNegativeSqrt struct {
	number float64
}

func (e *ErrNegativeSqrt) Error() string {
	newE := fmt.Sprint(float64(e.number))
	return fmt.Sprintf("Error at number: ", newE)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0.0, &ErrNegativeSqrt{x}
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
