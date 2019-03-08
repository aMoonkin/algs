package main

import "fmt"

type Printer func(contents string) (n int, err error)

func print2Std(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}

func main() {
	var p Printer
	p = print2Std
	p("somethins")
}
