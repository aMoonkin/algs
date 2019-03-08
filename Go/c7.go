package main

type operate func(x, y int) int

func calc(x, y int, op operate) (int, err) {
	if op == nil {
		return 0, error.new("invalid operation")
	}
	return op(x, y), nil
}
