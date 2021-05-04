package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldApplyToRangeIntegerInput(t *testing.T) {
	parser := &rangeNotationParser{}

	assert.Equal(t, true, parser.appliesTo("1-2"))

	assert.Equal(t, false, parser.appliesTo("1-2-3"))
	assert.Equal(t, false, parser.appliesTo("1.5"))
	assert.Equal(t, false, parser.appliesTo("*"))
	assert.Equal(t, false, parser.appliesTo("-1"))
	assert.Equal(t, false, parser.appliesTo("*/2"))
	assert.Equal(t, false, parser.appliesTo("3,4"))
	assert.Equal(t, false, parser.appliesTo("1"))
	assert.Equal(t, false, parser.appliesTo("text"))
}

func TestShouldReturnValuesWithinRange(t *testing.T) {
	parser := &rangeNotationParser{}
	input := "1-3"

	values, _ := parser.toValues(input, hour())

	assert.Equal(t, []int{1, 2, 3}, values)
}

func TestShouldReturnErrorIfRangeNotationIsInvalid(t *testing.T) {
	parser := &rangeNotationParser{}
	input := "-1-23"

	_, err := parser.toValues(input, hour())

	assert.Equal(t, "invalid range notation -1-23", err.Error())
}

func TestShouldReturnErrorIfInputIsNotRangeInput(t *testing.T) {
	parser := &rangeNotationParser{}
	input := "0"

	_, err := parser.toValues(input, hour())

	assert.Equal(t, "invalid range notation 0", err.Error())
}

func TestShouldReturnErrorIfRangeStartIsOutsideBoundsOfTimeUnit(t *testing.T) {
	parser := &rangeNotationParser{}
	input := "0-5"

	_, err := parser.toValues(input, month())

	assert.Equal(t, "invalid month value 0 outside bounds 1 and 12", err.Error())
}

func TestShouldReturnErrorIfRangeEndIsOutsideBoundsOfTimeUnit(t *testing.T) {
	parser := &rangeNotationParser{}
	input := "1-24"

	_, err := parser.toValues(input, hour())

	assert.Equal(t, "invalid hour value 24 outside bounds 0 and 23", err.Error())
}

func TestShouldReturnErrorIfRangeStartIsNotInteger(t *testing.T) {
	parser := &simpleNotationParser{}
	input := "1.5-20"

	_, err := parser.toValues(input, hour())

	assert.Equal(t, "strconv.Atoi: parsing \"1.5-20\": invalid syntax", err.Error())
}

func TestShouldReturnErrorIfRangeEndIsNotInteger(t *testing.T) {
	parser := &simpleNotationParser{}
	input := "1-20.5"

	_, err := parser.toValues(input, hour())

	assert.Equal(t, "strconv.Atoi: parsing \"1-20.5\": invalid syntax", err.Error())
}
