package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinuteShouldReturnAllValues(t *testing.T) {
	timeUnit := Minute()

	allValues := AllValues(timeUnit)

	assert.Equal(t, InclusiveRange(0, 59), allValues)
}

func TestMinuteShouldReturnValuesUnchanged(t *testing.T) {
	timeUnit := Minute()
	inputValues := "1,2,3"

	values := ToIntValues(timeUnit, inputValues)

	assert.Equal(t, inputValues, values)
}

func TestMinuteShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	timeUnit := Minute()
	values := []int{0, 59}

	err := ValidateMultiple(timeUnit, values)

	assert.Equal(t, nil, err)
}

func TestMinuteShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	timeUnit := Minute()
	value := -1

	err := Validate(timeUnit, value)

	assert.Equal(t, "invalid minute value -1 outside bounds 0 and 59", err.Error())
}

func TestMinuteShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	timeUnit := Minute()
	value := 60

	err := Validate(timeUnit, value)

	assert.Equal(t, "invalid minute value 60 outside bounds 0 and 59", err.Error())
}

func TestHourShouldReturnAllValues(t *testing.T) {
	timeUnit := Hour()

	allValues := AllValues(timeUnit)

	assert.Equal(t, InclusiveRange(0, 23), allValues)
}

func TestHourShouldReturnValuesUnchanged(t *testing.T) {
	timeUnit := Hour()
	inputValues := "1,2,3"

	values := ToIntValues(timeUnit, inputValues)

	assert.Equal(t, inputValues, values)
}

func TestHourShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	timeUnit := Hour()
	values := []int{0, 23}

	err := ValidateMultiple(timeUnit, values)

	assert.Equal(t, nil, err)
}

func TestHourShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	timeUnit := Hour()
	value := -1

	err := Validate(timeUnit, value)

	assert.Equal(t, "invalid hour value -1 outside bounds 0 and 23", err.Error())
}

func TestHourShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	timeUnit := Hour()
	value := 24

	err := Validate(timeUnit, value)

	assert.Equal(t, "invalid hour value 24 outside bounds 0 and 23", err.Error())
}

func TestDayOfMonthShouldReturnAllValues(t *testing.T) {
	timeUnit := DayOfMonth()

	allValues := AllValues(timeUnit)

	assert.Equal(t, InclusiveRange(1, 31), allValues)
}

func TestDayOfMonthShouldReturnValuesUnchanged(t *testing.T) {
	timeUnit := DayOfMonth()
	inputValues := "1,2,3"

	values := ToIntValues(timeUnit, inputValues)

	assert.Equal(t, inputValues, values)
}

func TestDayOfMonthShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	timeUnit := DayOfMonth()
	values := []int{1, 31}

	err := ValidateMultiple(timeUnit, values)

	assert.Equal(t, nil, err)
}

func TestDayOfMonthShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	timeUnit := DayOfMonth()
	value := 0

	err := Validate(timeUnit, value)

	assert.Equal(t, "invalid day of month value 0 outside bounds 1 and 31", err.Error())
}

func TestDayOfMonthShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	timeUnit := DayOfMonth()
	value := 32

	err := Validate(timeUnit, value)

	assert.Equal(t, "invalid day of month value 32 outside bounds 1 and 31", err.Error())
}

func TestMonthShouldReturnAllValues(t *testing.T) {
	timeUnit := Month()

	allValues := AllValues(timeUnit)

	assert.Equal(t, InclusiveRange(1, 12), allValues)
}

func TestMonthShouldReturnTextualValuesConvertedToInts(t *testing.T) {
	timeUnit := Month()
	inputValues := "JAN,Feb,maR,aPr,may,jun,jul,aug,sep,oct,nov,dec"

	values := ToIntValues(timeUnit, inputValues)

	assert.Equal(t, "1,2,3,4,5,6,7,8,9,10,11,12", values)
}

func TestMonthShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	timeUnit := Month()
	values := []int{1, 12}

	err := ValidateMultiple(timeUnit, values)

	assert.Equal(t, nil, err)
}

func TestMonthShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	timeUnit := Month()
	value := 0

	err := Validate(timeUnit, value)

	assert.Equal(t, "invalid month value 0 outside bounds 1 and 12", err.Error())
}

func TestMonthShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	timeUnit := Month()
	value := 13

	err := Validate(timeUnit, value)

	assert.Equal(t, "invalid month value 13 outside bounds 1 and 12", err.Error())
}

func TestDayOfWeekShouldReturnValues(t *testing.T) {
	timeUnit := DayOfWeek()

	allValues := AllValues(timeUnit)

	assert.Equal(t, InclusiveRange(0, 6), allValues)
}

func TestDayOfWeekShouldReturnTextualValuesConvertedToInts(t *testing.T) {
	timeUnit := DayOfWeek()
	inputValues := "MON,Tue,weD,tHu,fri,sat,sun"

	values := ToIntValues(timeUnit, inputValues)

	assert.Equal(t, "0,1,2,3,4,5,6", values)
}

func TestDayOfWeekShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	timeUnit := DayOfWeek()
	values := []int{0, 6}

	err := ValidateMultiple(timeUnit, values)

	assert.Equal(t, nil, err)
}

func TestDayOfWeekShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	timeUnit := DayOfWeek()
	value := -1

	err := Validate(timeUnit, value)

	assert.Equal(t, "invalid day of week value -1 outside bounds 0 and 6", err.Error())
}

func TestDayOfWeekShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	timeUnit := DayOfWeek()
	value := 7

	err := Validate(timeUnit, value)

	assert.Equal(t, "invalid day of week value 7 outside bounds 0 and 6", err.Error())
}
