package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldApplyToWildcard(t *testing.T) {
	parser := &wildcardNotationParser{}

	assert.Equal(t, true, parser.appliesTo("*"))

	assert.Equal(t, false, parser.appliesTo("1.5"))
	assert.Equal(t, false, parser.appliesTo("1"))
	assert.Equal(t, false, parser.appliesTo("-1"))
	assert.Equal(t, false, parser.appliesTo("*/2"))
	assert.Equal(t, false, parser.appliesTo("3,4"))
	assert.Equal(t, false, parser.appliesTo("5-6"))
	assert.Equal(t, false, parser.appliesTo("text"))
}

func TestShouldReturnAllValuesFromTimeUnitRegardlessOfInput(t *testing.T) {
	parser := &wildcardNotationParser{}
	input := "*"
	timeUnit := Hour()

	values, _ := parser.toValues(input, timeUnit)

	assert.Equal(t, AllValues(timeUnit), values)
}
