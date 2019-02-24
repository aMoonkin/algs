package BigO

func cal2(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum += i * j
		}
	}
	return sum
}
