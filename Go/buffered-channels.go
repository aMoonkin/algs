package main

import "fmt"

var maxNum = 2

func main() {
	ch := make(chan int, maxNum)

	for index := 0; index < maxNum; index++ {
		ch <- index + 1
	}

	for _, v := range ch {
		fmt.Println(v)
	}
}
