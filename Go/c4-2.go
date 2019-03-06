package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}

	_, ok := interface{}(container).([]string)
	_, ok2 := interface{}(container).(map[int]string)

	if !(ok || ok2) {
		fmt.Println("err")
	}

	elem, err := getElement(container)
	if err != nil {
		return
	}
	fmt.Println(elem)
}

func getElement(containerI interface{}) (elem string, err error) {
	switch t := containerI.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		elem = t[1]
	default:
		return
	}
	return
}
