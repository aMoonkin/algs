package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	values := list.New()
	values.PushBack("bird")
	values.PushBack("cat")
	values.PushBack("snake")

	for index := 0; index < 20; index++ {
		values.PushFront(strconv.Itoa(index))
	}

	for temp := values.Front(); temp != nil; temp = temp.Next() {
		fmt.Println(temp.Value)
	}
}
