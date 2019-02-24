package BigO

var arr = make([]int, 10)
var i = 0

func add(element int) {
	if i >= len(arr) {
		new_arr := make([]int, 2*len(arr))

		for i, v := range arr {
			new_arr[v] = i
		}

		arr = new_arr
	}
	arr[i] = element
	i++
}
