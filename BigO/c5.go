package BigO

var n = 1
var array = make([]int, n)
var count = 0

func insert(val int) {
	if count == len(array) {
		sum := 0
		for _, i := range array {
			sum += i
		}
		array[0] = sum
		count = 1
	}

	array[count] = val
	count++
}
