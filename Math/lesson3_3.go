package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []string{"s", "t", "y", "a"}
	sort.Strings(arr)
	fmt.Println(findWord(arr, "s"))
}

func findWord(dict []string, wordToFind string) int {
	if len(dict) == 0 {
		return -1
	}

	x := 0
	y := len(dict) - 1

	for x < y {
		m := (y - x) / 2
		if dict[m] == wordToFind {
			return m
		} else if dict[m] < wordToFind {
			x = m
		} else {
			y = m
		}
	}

	return x

}
