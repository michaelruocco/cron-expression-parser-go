package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnErrorIfEmptyArgumentsPassed(t *testing.T) {
	args := []string{}

	_, err := sanitize(args)

	assert.Equal(t, err.Error(), "usage: please provide a valid cron expression")
}

func TestShouldReturnErrorIfLessThanSixArgumentsPassed(t *testing.T) {
	args := []string{"1", "2", "3", "4", "5"}

	_, err := sanitize(args)

	assert.Equal(t, err.Error(), "usage: please provide a valid cron expression")
}

func TestShouldReturnAllArgumentsIfSixArgumentsPassed(t *testing.T) {
	args := []string{"1", "2", "3", "4", "5", "6"}

	sanitized, _ := sanitize(args)

	assert.Equal(t, sanitized, args)
}

func TestShouldReturnLastSixArgumentsIfMoreThanSixArgumentsPassed(t *testing.T) {
	args := []string{"1", "2", "3", "4", "5", "6", "7"}

	sanitized, _ := sanitize(args)

	assert.Equal(t, []string{"2", "3", "4", "5", "6", "7"}, sanitized)
}
