package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = getTheFlag()
	flag.Parse()
	fmt.Println(*name)
}

func getTheFlag() *string {
	return flag.String("name", "everyone", "greeting ")
}
