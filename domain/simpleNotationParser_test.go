package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldApplyToPlainIntegerInput(t *testing.T) {
	parser := &simpleNotationParser{}

	assert.Equal(t, true, parser.appliesTo("1"))

	assert.Equal(t, false, parser.appliesTo("1.5"))
	assert.Equal(t, false, parser.appliesTo("*"))
	assert.Equal(t, false, parser.appliesTo("-1"))
	assert.Equal(t, false, parser.appliesTo("*/2"))
	assert.Equal(t, false, parser.appliesTo("3,4"))
	assert.Equal(t, false, parser.appliesTo("5-6"))
	assert.Equal(t, false, parser.appliesTo("text"))
}

func TestShouldReturnIntegerForValidIntegerInput(t *testing.T) {
	parser := &simpleNotationParser{}
	input := "1"

	values, _ := parser.toValues(input, Hour())

	assert.Equal(t, []int{1}, values)
}

func TestShouldReturnErrorIfInputOutsideBoundsOfTimeUnit(t *testing.T) {
	parser := &simpleNotationParser{}
	input := "24"

	_, err := parser.toValues(input, Hour())

	assert.Equal(t, "invalid hour value 24 outside bounds 0 and 23", err.Error())
}

func TestShouldReturnErrorIfInputIsNotInteger(t *testing.T) {
	parser := &simpleNotationParser{}
	input := "2.5"

	_, err := parser.toValues(input, Hour())

	assert.Equal(t, "strconv.Atoi: parsing \"2.5\": invalid syntax", err.Error())
}
