package main

import (
	"errors"
	"fmt"
)

func echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("empty req")
		return
	}
	response = fmt.Sprintf("echo: %s", request)
	return
}

func main() {
	for _, req := range []string{"", "hello!"} {
		fmt.Printf("request: %s\n", req)
		resp, err := echo(req)
		if err != nil {
			fmt.Printf("err %s", err)
			continue
		}
		fmt.Printf("response: %s\n", resp)
	}
}
