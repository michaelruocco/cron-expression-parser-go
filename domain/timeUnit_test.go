package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnAllMinuteValues(t *testing.T) {
	timeUnit := Minute()

	allValues := AllValues(timeUnit)

	assert.Equal(t, allValues, allRangeValuesIncluding(0, 59))
}

func TestShouldReturnAllHourValues(t *testing.T) {
	timeUnit := Hour()

	allValues := AllValues(timeUnit)

	assert.Equal(t, allValues, allRangeValuesIncluding(0, 23))
}

func TestShouldReturnAllDayOfMonthValues(t *testing.T) {
	timeUnit := DayOfMonth()

	allValues := AllValues(timeUnit)

	assert.Equal(t, allValues, allRangeValuesIncluding(1, 31))
}

func TestShouldReturnAllMonthValues(t *testing.T) {
	timeUnit := Month()

	allValues := AllValues(timeUnit)

	assert.Equal(t, allValues, allRangeValuesIncluding(1, 12))
}

func TestShouldReturnAllDayOfWeekValues(t *testing.T) {
	timeUnit := DayOfWeek()

	allValues := AllValues(timeUnit)

	assert.Equal(t, allValues, allRangeValuesIncluding(0, 6))
}

func allRangeValuesIncluding(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
