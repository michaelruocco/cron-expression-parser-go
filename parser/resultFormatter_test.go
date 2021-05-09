package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldFormatResult(t *testing.T) {
	result := buildResult()

	formatted := format(result)

	assert.Equal(t, "minute        39 40\n"+
		"hour          9 10\n"+
		"day of month  20 21\n"+
		"month         3 4\n"+
		"day of week   1 2\n"+
		"command       my-command", formatted)
}

func buildResult() cronResult {
	return cronResult{
		Minutes:     []int{40, 39},
		Hours:       []int{10, 9},
		DaysOfMonth: []int{21, 20},
		Months:      []int{4, 3},
		DaysOfWeek:  []int{2, 1},
		Command:     "my-command",
	}
}
