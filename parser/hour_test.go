package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHourShouldReturnAllValues(t *testing.T) {
	hour := hour()

	allValues := hour.allValues()

	assert.Equal(t, inclusiveRange(0, 23), allValues)
}

func TestHourShouldReturnValuesUnchanged(t *testing.T) {
	hour := hour()
	inputValues := "1,2,3"

	values := hour.toIntValues(inputValues)

	assert.Equal(t, inputValues, values)
}

func TestHourShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	hour := hour()
	value := []int{0, 23}

	err := hour.validate(value)

	assert.Equal(t, nil, err)
}

func TestHourShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	hour := hour()
	value := []int{-1}

	err := hour.validate(value)

	assert.Equal(t, "invalid hour value -1 outside bounds 0 and 23", err.Error())
}

func TestHourShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	hour := hour()
	value := []int{24}

	err := hour.validate(value)

	assert.Equal(t, "invalid hour value 24 outside bounds 0 and 23", err.Error())
}
