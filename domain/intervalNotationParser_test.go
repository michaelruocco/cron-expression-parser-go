package domain

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

	assert.Equal(t, []int{2, 3, 4, 8, 9, 10, 14, 15, 16, 20, 21, 22}, values)
}

func TestShouldReturnIntervalValuesWithMultipleStartValues(t *testing.T) {
	parser := buildIntervalNotationParser()
	input := "2,4/6"

	values, _ := parser.toValues(input, hour())

	assert.Equal(t, []int{2, 4, 8, 10, 14, 16, 20, 22}, values)
}

func TestShouldReturnIntervalValuesWithMultipleStartValues1(t *testing.T) {
	parser := buildIntervalNotationParser()
	input := "3,45/15"

	values, _ := parser.toValues(input, minute())

	assert.Equal(t, []int{3, 18, 33, 45, 48}, values)
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
