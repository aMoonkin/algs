package main

import (
	"errors"
	"fmt"
)

type operator func(x, y int) int

type calculatorFunc func(a, b int) (int, error)

func getCalculator(op operator) calculatorFunc {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("type")
		}
		return op(x, y), nil
	}
}

func main() {
	x, y := 12, 23
	op := func(x, y int) int {
		return x + y
	}
	add := getCalculator(op)
	res, err := add(x, y)
	fmt.Println(res, err)
}
