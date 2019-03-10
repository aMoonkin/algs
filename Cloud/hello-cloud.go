package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, cloud")
	fmt.Println("你好，云")
	for index := 0; index < 1000; index++ {
		fmt.Println(dieRoll(3))
	}
}

func dieRoll(size int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(size) + 1
}
