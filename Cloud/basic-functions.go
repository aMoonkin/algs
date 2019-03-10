package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println(dieRoll(6))
	res1, res2 := rollTwo(6, 10)
	fmt.Println(res1, res2)
	named, err := returnsNamed("global", 42)
	fmt.Println(named, err)

}

func dieRoll(size int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(size) + 1
}

func rollTwo(size1, size2 int) (int, int) {
	return dieRoll(size1), dieRoll(size2)
}

func returnsNamed(input1 string, input2 int) (theResult string, err error) {
	theResult = "modified " + input1 + ", " + strconv.Itoa(input2)
	return theResult, err
}
