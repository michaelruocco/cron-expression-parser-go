package parser

func inclusiveRange(x int, y int) []int {
	start := min(x, y)
	end := max(x, y)
	values := make([]int, end-start+1)
	for i := range values {
		values[i] = start + i
	}
	return values
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
