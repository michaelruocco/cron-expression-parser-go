package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnTrueIfValueIsPositiveInteger(t *testing.T) {
	assert.Equal(t, true, IsInt("1"))
	assert.Equal(t, true, IsInt("0"))

	assert.Equal(t, false, IsInt("-1"))
	assert.Equal(t, false, IsInt("1.1"))
	assert.Equal(t, false, IsInt("text"))
}
