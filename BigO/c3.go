package BigO

import "fmt"

func print(n int) {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i * i
	}

	for i := n - 1; i >= 0; i-- {
		fmt.Println(a[i])
	}
}
