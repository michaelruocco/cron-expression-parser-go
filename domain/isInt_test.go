package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnTrueIfValueIsPositiveInteger(t *testing.T) {
	assert.Equal(t, true, isInt("1"))
	assert.Equal(t, true, isInt("0"))

	assert.Equal(t, false, isInt("-1"))
	assert.Equal(t, false, isInt("1.1"))
	assert.Equal(t, false, isInt("text"))
}
