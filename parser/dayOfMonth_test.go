package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayOfMonthShouldReturnAllValues(t *testing.T) {
	dayOfMonth := dayOfMonth()

	allValues := dayOfMonth.allValues()

	assert.Equal(t, inclusiveRange(1, 31), allValues)
}

func TestDayOfMonthShouldReturnValuesUnchanged(t *testing.T) {
	dayOfMonth := dayOfMonth()
	inputValues := "1,2,3"

	values := dayOfMonth.toIntValues(inputValues)

	assert.Equal(t, inputValues, values)
}

func TestDayOfMonthShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	dayOfMonth := dayOfMonth()
	values := []int{1, 31}

	err := dayOfMonth.validate(values)

	assert.Equal(t, nil, err)
}

func TestDayOfMonthShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	dayOfMonth := dayOfMonth()
	values := []int{0}

	err := dayOfMonth.validate(values)

	assert.Equal(t, "invalid day of month value 0 outside bounds 1 and 31", err.Error())
}

func TestDayOfMonthShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	dayOfMonth := dayOfMonth()
	values := []int{32}

	err := dayOfMonth.validate(values)

	assert.Equal(t, "invalid day of month value 32 outside bounds 1 and 31", err.Error())
}
