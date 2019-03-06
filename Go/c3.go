package main

import "fmt"

var c string = "1"

func main() {
	var c string = "jajaaj"

	{
		var c = 555444
		fmt.Println(c)
	}
	c = "555"
	fmt.Println(c)
}
