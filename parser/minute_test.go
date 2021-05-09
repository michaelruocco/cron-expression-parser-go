package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinuteShouldReturnAllValues(t *testing.T) {
	minute := minute()

	allValues := minute.allValues()

	assert.Equal(t, inclusiveRange(0, 59), allValues)
}

func TestMinuteShouldReturnValuesUnchanged(t *testing.T) {
	minute := minute()
	inputValues := "1,2,3"

	values := minute.toIntValues(inputValues)

	assert.Equal(t, inputValues, values)
}

func TestMinuteShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	minute := minute()
	values := []int{0, 59}

	err := minute.validate(values)

	assert.Equal(t, nil, err)
}

func TestMinuteShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	minute := minute()
	values := []int{-1}

	err := minute.validate(values)

	assert.Equal(t, "invalid minute value -1 outside bounds 0 and 59", err.Error())
}

func TestMinuteShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	minute := minute()
	values := []int{60}

	err := minute.validate(values)

	assert.Equal(t, "invalid minute value 60 outside bounds 0 and 59", err.Error())
}
