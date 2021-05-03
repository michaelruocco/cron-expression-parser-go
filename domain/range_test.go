package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnValuesInRangeIncludingStartAndEnd(t *testing.T) {
	start := 0
	end := 3

	values := InclusiveRange(start, end)

	assert.Equal(t, []int{0, 1, 2, 3}, values)
}
