package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnValuesInRangeIncludingStartAndEnd(t *testing.T) {
	start := 0
	end := 3

	values := inclusiveRange(start, end)

	assert.Equal(t, []int{0, 1, 2, 3}, values)
}

func TestShouldReturnValuesInRangeIncludingStartAndEndWhenStartGreaterThanEnd(t *testing.T) {
	start := 3
	end := 0

	values := inclusiveRange(start, end)

	assert.Equal(t, []int{0, 1, 2, 3}, values)
}

func TestShouldReturnValueWhenRangeStartAndEndAreEqual(t *testing.T) {
	start := 3
	end := 3

	values := inclusiveRange(start, end)

	assert.Equal(t, []int{3}, values)
}
