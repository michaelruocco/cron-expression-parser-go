package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayOfWeekShouldReturnValues(t *testing.T) {
	dayOfWeek := dayOfWeek()

	allValues := dayOfWeek.allValues()

	assert.Equal(t, inclusiveRange(0, 6), allValues)
}

func TestDayOfWeekShouldReturnTextualValuesConvertedToInts(t *testing.T) {
	dayOfWeek := dayOfWeek()
	inputValues := "MON,Tue,weD,tHu,fri,sat,sun"

	values := dayOfWeek.toIntValues(inputValues)

	assert.Equal(t, "0,1,2,3,4,5,6", values)
}

func TestDayOfWeekShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	dayOfWeek := dayOfWeek()
	values := []int{0, 6}

	err := dayOfWeek.validate(values...)

	assert.Equal(t, nil, err)
}

func TestDayOfWeekShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	dayOfWeek := dayOfWeek()
	values := -1

	err := dayOfWeek.validate(values)

	assert.Equal(t, "invalid day of week value -1 outside bounds 0 and 6", err.Error())
}

func TestDayOfWeekShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	dayOfWeek := dayOfWeek()
	values := 7

	err := dayOfWeek.validate(values)

	assert.Equal(t, "invalid day of week value 7 outside bounds 0 and 6", err.Error())
}
