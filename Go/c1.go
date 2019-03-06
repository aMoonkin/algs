package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
	var name2 = *flag.String("name1", "hhaha", "nothing")
	flag.Parse()
	fmt.Println("Hello", name, name2)
}
