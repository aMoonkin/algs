package BigO

func find(array []int, x int) int {
	pos := -1
	for _, i := range array {
		if array[i] == x {
			pos = i
			break
		}
	}
	return pos
}
