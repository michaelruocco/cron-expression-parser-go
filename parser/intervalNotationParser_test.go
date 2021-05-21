package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldApplyToIntervalInputWithRangeOrMultipleStartValues(t *testing.T) {
	parser := buildIntervalNotationParser()

	assert.Equal(t, true, parser.appliesTo("2/2"))
	assert.Equal(t, true, parser.appliesTo("*/6"))
	assert.Equal(t, true, parser.appliesTo("2-4/6"))
	assert.Equal(t, true, parser.appliesTo("2,4/6"))

	assert.Equal(t, false, parser.appliesTo("2-4.4/6"))
	assert.Equal(t, false, parser.appliesTo("2.2,4/6"))
	assert.Equal(t, false, parser.appliesTo("1-2-3"))
	assert.Equal(t, false, parser.appliesTo("1.5"))
	assert.Equal(t, false, parser.appliesTo("*"))
	assert.Equal(t, false, parser.appliesTo("-1"))
	assert.Equal(t, false, parser.appliesTo("3,4"))
	assert.Equal(t, false, parser.appliesTo("1"))
	assert.Equal(t, false, parser.appliesTo("text"))
}

func TestShouldReturnIntervalValues(t *testing.T) {
	parser := buildIntervalNotationParser()
	input := "2/3"

	values, _ := parser.toValues(input, hour())

	assert.Equal(t, []int{2, 5, 8, 11, 14, 17, 20, 23}, values)
}

func TestShouldReturnIntervalValuesWithRangeOfStartValues(t *testing.T) {
	parser := buildIntervalNotationParser()
	input := "2-4/6"

	values, _ := parser.toValues(input, hour())

	assert.Equal(t, []int{2}, values)
}

func TestShouldReturnIntervalValuesWithStartAndEndValues(t *testing.T) {
	parser := buildIntervalNotationParser()
	input := "2,4/6"

	values, _ := parser.toValues(input, hour())

	assert.Equal(t, []int{2, 8, 14, 20}, values)
}

func TestShouldReturnIntervalValuesWithNonIntegerEndValue(t *testing.T) {
	parser := buildIntervalNotationParser()
	input := "2,20.5/6"

	_, err := parser.toValues(input, hour())

	assert.Equal(t, "invalid notation 2,20.5", err.Error())
}

func TestShouldReturnIntervalValuesWithEndValueOutOfBoundsValue(t *testing.T) {
	parser := buildIntervalNotationParser()
	input := "2,24/6"

	_, err := parser.toValues(input, hour())

	assert.Equal(t, "invalid hour value 24 outside bounds 0 and 23", err.Error())
}

func TestShouldReturnIntervalValuesWithWildcardStartValues(t *testing.T) {
	parser := buildIntervalNotationParser()
	input := "*/6"

	values, _ := parser.toValues(input, hour())

	assert.Equal(t, []int{0, 6, 12, 18}, values)
}

func TestShouldReturnErrorIfInputIsNotInterval(t *testing.T) {
	parser := buildIntervalNotationParser()
	input := "2"

	_, err := parser.toValues(input, hour())

	assert.Equal(t, "invalid interval notation 2", err.Error())
}

func TestShouldReturnErrorIfStartIsOutsideBoundsOfTimeUnit(t *testing.T) {
	parser := buildIntervalNotationParser()
	input := "24/2"

	_, err := parser.toValues(input, hour())

	assert.Equal(t, "invalid hour value 24 outside bounds 0 and 23", err.Error())
}

func TestShouldReturnErrorIfIntervalStartIsIsNotInteger(t *testing.T) {
	parser := buildIntervalNotationParser()
	input := "3.5/2"

	_, err := parser.toValues(input, hour())

	assert.Equal(t, "invalid notation 3.5", err.Error())
}

func TestShouldReturnErrorIfIntervalIsNotInteger(t *testing.T) {
	parser := buildIntervalNotationParser()
	input := "3/2.5"

	_, err := parser.toValues(input, hour())

	assert.Equal(t, "invalid interval notation 3/2.5", err.Error())
}
