package main

import (
	"fmt"
	"math/rand"
	"time"
)

func dieRoll(size int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(size) + 1
}

type dieRollFunc func(int) int

func fakeDieRoll(sz int) int {
	return 42
}

func getDieRolls() []dieRollFunc {
	return []dieRollFunc{
		dieRoll,
		fakeDieRoll,
	}
}

func main() {
	for _, rollFunc := range getDieRolls() {
		fmt.Println(rollFunc(10))
	}
}
