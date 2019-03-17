package main

import "fmt"

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for index := 0; index < n; index++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
