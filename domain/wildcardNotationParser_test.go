package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldApplyToWildcard(t *testing.T) {
	parser := &wildcardNotationParser{}

	assert.Equal(t, parser.appliesTo("*"), true)

	assert.Equal(t, parser.appliesTo("1.5"), false)
	assert.Equal(t, parser.appliesTo("1"), false)
	assert.Equal(t, parser.appliesTo("-1"), false)
	assert.Equal(t, parser.appliesTo("*/2"), false)
	assert.Equal(t, parser.appliesTo("3,4"), false)
	assert.Equal(t, parser.appliesTo("5-6"), false)
	assert.Equal(t, parser.appliesTo("text"), false)
}

func TestShouldReturnAllValuesFromTimeUnitRegardlessOfInput(t *testing.T) {
	parser := &wildcardNotationParser{}
	input := "*"
	timeUnit := Hour()

	values, _ := parser.toValues(input, timeUnit)

	assert.Equal(t, values, AllValues(timeUnit))
}
