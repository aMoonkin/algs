package main

import "fmt"

func main() {
	a1 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("a1: %v, (len: %d, cap: %d) \n",
		a1, len(a1), cap(a1))

	s9 := a1[1:4]

	fmt.Printf("s9: %v (len: %d, cap: %d)\n",
		s9, len(s9), cap(s9))

	for index := 0; index < 5; index++ {
		s9 = append(s9, index+999)
		fmt.Printf("s9: %v (len: %d, cap: %d)\n",
			s9, len(s9), cap(s9))
	}

	fmt.Printf("a1: %v (len: %d, cap: %d)\n",
		a1, len(a1), cap(a1))
	fmt.Println()
}
