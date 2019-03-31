package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getSquareRoot(2))
}

func getSquareRoot(n float64) float64 {
	x := 1.0
	y := n
	for y-x > 0.001 {
		m := 0.5 * (y + x)
		if math.Pow(m, 2) > n {
			y = m
		} else {
			x = m
		}
	}
	return x
}
