package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldApplyToPlainIntegerInput(t *testing.T) {
	parser := &simpleNotationParser{}

	assert.Equal(t, parser.appliesTo("1"), true)

	assert.Equal(t, parser.appliesTo("1.5"), false)
	assert.Equal(t, parser.appliesTo("*"), false)
	assert.Equal(t, parser.appliesTo("-1"), false)
	assert.Equal(t, parser.appliesTo("*/2"), false)
	assert.Equal(t, parser.appliesTo("3,4"), false)
	assert.Equal(t, parser.appliesTo("5-6"), false)
	assert.Equal(t, parser.appliesTo("text"), false)
}

func TestShouldReturnIntegerForValidIntegerInput(t *testing.T) {
	parser := &simpleNotationParser{}
	input := "1"

	values, _ := parser.toValues(input, Hour())

	assert.Equal(t, values, []int{1})
}

func TestShouldReturnErrorIfInputOutsideBoundsOfTimeUnit(t *testing.T) {
	parser := &simpleNotationParser{}
	input := "24"

	_, err := parser.toValues(input, Hour())

	assert.Equal(t, err.Error(), "invalid hour value 24 outside bounds 0 and 23")
}

func TestShouldReturnErrorIfInputIsNotInteger(t *testing.T) {
	parser := &simpleNotationParser{}
	input := "2.5"

	_, err := parser.toValues(input, Hour())

	assert.Equal(t, err.Error(), "strconv.Atoi: parsing \"2.5\": invalid syntax")
}
