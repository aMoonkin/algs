package main

import "fmt"

func main() {
	result := make([]int, 0)
	lesson5_1(10, result)
}

func lesson5_1(totalReword int, result []int) {

	rewards := []int{1, 2, 5, 10}

	if totalReword < 0 {
		return
	}

	if totalReword == 0 {
		fmt.Println(result)
	}

	for _, reward := range rewards {
		newResult := append(result, reward)
		lesson5_1(totalReword-reward, newResult)
	}

}
