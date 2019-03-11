package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Println(name)
}
