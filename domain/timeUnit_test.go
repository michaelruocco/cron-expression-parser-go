package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinuteShouldReturnAllValues(t *testing.T) {
	timeUnit := Minute()

	allValues := AllValues(timeUnit)

	assert.Equal(t, allValues, allRangeValuesIncluding(0, 59))
}

func TestMinuteShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	timeUnit := Minute()
	values := []int{0, 59}

	err := ValidateMultiple(timeUnit, values)

	assert.Equal(t, err, nil)
}

func TestMinuteShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	timeUnit := Minute()
	value := -1

	err := Validate(timeUnit, value)

	assert.Equal(t, err.Error(), "invalid minute value -1 outside bounds 0 and 59")
}

func TestMinuteShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	timeUnit := Minute()
	value := 60

	err := Validate(timeUnit, value)

	assert.Equal(t, err.Error(), "invalid minute value 60 outside bounds 0 and 59")
}

func TestHourShouldReturnAllValues(t *testing.T) {
	timeUnit := Hour()

	allValues := AllValues(timeUnit)

	assert.Equal(t, allValues, allRangeValuesIncluding(0, 23))
}

func TestHourShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	timeUnit := Hour()
	values := []int{0, 23}

	err := ValidateMultiple(timeUnit, values)

	assert.Equal(t, err, nil)
}

func TestHourShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	timeUnit := Hour()
	value := -1

	err := Validate(timeUnit, value)

	assert.Equal(t, err.Error(), "invalid hour value -1 outside bounds 0 and 23")
}

func TestHourShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	timeUnit := Hour()
	value := 24

	err := Validate(timeUnit, value)

	assert.Equal(t, err.Error(), "invalid hour value 24 outside bounds 0 and 23")
}

func TestDayOfMonthShouldReturnAllValues(t *testing.T) {
	timeUnit := DayOfMonth()

	allValues := AllValues(timeUnit)

	assert.Equal(t, allValues, allRangeValuesIncluding(1, 31))
}

func TestDayOfMonthShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	timeUnit := DayOfMonth()
	values := []int{1, 31}

	err := ValidateMultiple(timeUnit, values)

	assert.Equal(t, err, nil)
}

func TestDayOfMonthShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	timeUnit := DayOfMonth()
	value := 0

	err := Validate(timeUnit, value)

	assert.Equal(t, err.Error(), "invalid day of month value 0 outside bounds 1 and 31")
}

func TestDayOfMonthShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	timeUnit := DayOfMonth()
	value := 32

	err := Validate(timeUnit, value)

	assert.Equal(t, err.Error(), "invalid day of month value 32 outside bounds 1 and 31")
}

func TestMonthShouldReturnAllValues(t *testing.T) {
	timeUnit := Month()

	allValues := AllValues(timeUnit)

	assert.Equal(t, allValues, allRangeValuesIncluding(1, 12))
}

func TestMonthShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	timeUnit := Month()
	values := []int{1, 12}

	err := ValidateMultiple(timeUnit, values)

	assert.Equal(t, err, nil)
}

func TestMonthShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	timeUnit := Month()
	value := 0

	err := Validate(timeUnit, value)

	assert.Equal(t, err.Error(), "invalid month value 0 outside bounds 1 and 12")
}

func TestMonthShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	timeUnit := Month()
	value := 13

	err := Validate(timeUnit, value)

	assert.Equal(t, err.Error(), "invalid month value 13 outside bounds 1 and 12")
}

func TestDayOfWeekShouldReturnValues(t *testing.T) {
	timeUnit := DayOfWeek()

	allValues := AllValues(timeUnit)

	assert.Equal(t, allValues, allRangeValuesIncluding(0, 6))
}

func TestDayOfWeekShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	timeUnit := DayOfWeek()
	values := []int{0, 6}

	err := ValidateMultiple(timeUnit, values)

	assert.Equal(t, err, nil)
}

func TestDayOfWeekShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	timeUnit := DayOfWeek()
	value := -1

	err := Validate(timeUnit, value)

	assert.Equal(t, err.Error(), "invalid day of week value -1 outside bounds 0 and 6")
}

func TestDayOfWeekShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	timeUnit := DayOfWeek()
	value := 7

	err := Validate(timeUnit, value)

	assert.Equal(t, err.Error(), "invalid day of week value 7 outside bounds 0 and 6")
}

func allRangeValuesIncluding(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
