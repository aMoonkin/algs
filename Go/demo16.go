package main

import "fmt"

func main() {
	s6 := make([]int, 0)
	fmt.Printf("cap is %d\n", cap(s6))

	for index := 0; index < 5; index++ {
		s6 = append(s6, index+1)
		fmt.Printf("s6(%d): len: %d, cap: %d\n", index, len(s6), cap(s6))
	}

	s7 := make([]int, 1024)
	fmt.Printf("The capacity of s7: %d\n", cap(s7))

	s7e2 := append(s7, make([]int, 400)...)
	fmt.Printf("s7e2: len: %d, cap: %d\n", len(s7e2), cap(s7e2))
	s7e3 := append(s7, make([]int, 600)...)
	fmt.Printf("s7e3: len: %d, cap: %d\n", len(s7e3), cap(s7e3))
	fmt.Println()
}
