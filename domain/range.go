package domain

func inclusiveRange(start int, end int) []int {
	values := make([]int, end-start+1)
	for i := range values {
		values[i] = start + i
	}
	return values
}
