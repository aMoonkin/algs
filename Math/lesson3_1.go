package main

import "fmt"

func lesson3_1(grid int) int {
	sum := 0
	firstGrid := 1
	sum += firstGrid

	for index := 2; index <= grid; index++ {
		firstGrid *= 2
		sum += firstGrid
	}
	return sum
}

func main() {
	fmt.Println(lesson3_1(60))
}
