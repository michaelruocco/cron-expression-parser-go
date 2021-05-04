package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinuteShouldReturnAllValues(t *testing.T) {
	timeUnit := minute()

	allValues := timeUnitToAllValues(timeUnit)

	assert.Equal(t, inclusiveRange(0, 59), allValues)
}

func TestMinuteShouldReturnValuesUnchanged(t *testing.T) {
	timeUnit := minute()
	inputValues := "1,2,3"

	values := timeUnitInputToIntValues(timeUnit, inputValues)

	assert.Equal(t, inputValues, values)
}

func TestMinuteShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	timeUnit := minute()
	values := []int{0, 59}

	err := timeUnitValidateMultipleInputs(timeUnit, values)

	assert.Equal(t, nil, err)
}

func TestMinuteShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	timeUnit := minute()
	value := -1

	err := timeUnitValidateInput(timeUnit, value)

	assert.Equal(t, "invalid minute value -1 outside bounds 0 and 59", err.Error())
}

func TestMinuteShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	timeUnit := minute()
	value := 60

	err := timeUnitValidateInput(timeUnit, value)

	assert.Equal(t, "invalid minute value 60 outside bounds 0 and 59", err.Error())
}

func TestHourShouldReturnAllValues(t *testing.T) {
	timeUnit := hour()

	allValues := timeUnitToAllValues(timeUnit)

	assert.Equal(t, inclusiveRange(0, 23), allValues)
}

func TestHourShouldReturnValuesUnchanged(t *testing.T) {
	timeUnit := hour()
	inputValues := "1,2,3"

	values := timeUnitInputToIntValues(timeUnit, inputValues)

	assert.Equal(t, inputValues, values)
}

func TestHourShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	timeUnit := hour()
	values := []int{0, 23}

	err := timeUnitValidateMultipleInputs(timeUnit, values)

	assert.Equal(t, nil, err)
}

func TestHourShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	timeUnit := hour()
	value := -1

	err := timeUnitValidateInput(timeUnit, value)

	assert.Equal(t, "invalid hour value -1 outside bounds 0 and 23", err.Error())
}

func TestHourShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	timeUnit := hour()
	value := 24

	err := timeUnitValidateInput(timeUnit, value)

	assert.Equal(t, "invalid hour value 24 outside bounds 0 and 23", err.Error())
}

func TestDayOfMonthShouldReturnAllValues(t *testing.T) {
	timeUnit := dayOfMonth()

	allValues := timeUnitToAllValues(timeUnit)

	assert.Equal(t, inclusiveRange(1, 31), allValues)
}

func TestDayOfMonthShouldReturnValuesUnchanged(t *testing.T) {
	timeUnit := dayOfMonth()
	inputValues := "1,2,3"

	values := timeUnitInputToIntValues(timeUnit, inputValues)

	assert.Equal(t, inputValues, values)
}

func TestDayOfMonthShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	timeUnit := dayOfMonth()
	values := []int{1, 31}

	err := timeUnitValidateMultipleInputs(timeUnit, values)

	assert.Equal(t, nil, err)
}

func TestDayOfMonthShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	timeUnit := dayOfMonth()
	value := 0

	err := timeUnitValidateInput(timeUnit, value)

	assert.Equal(t, "invalid day of month value 0 outside bounds 1 and 31", err.Error())
}

func TestDayOfMonthShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	timeUnit := dayOfMonth()
	value := 32

	err := timeUnitValidateInput(timeUnit, value)

	assert.Equal(t, "invalid day of month value 32 outside bounds 1 and 31", err.Error())
}

func TestMonthShouldReturnAllValues(t *testing.T) {
	timeUnit := month()

	allValues := timeUnitToAllValues(timeUnit)

	assert.Equal(t, inclusiveRange(1, 12), allValues)
}

func TestMonthShouldReturnTextualValuesConvertedToInts(t *testing.T) {
	timeUnit := month()
	inputValues := "JAN,Feb,maR,aPr,may,jun,jul,aug,sep,oct,nov,dec"

	values := timeUnitInputToIntValues(timeUnit, inputValues)

	assert.Equal(t, "1,2,3,4,5,6,7,8,9,10,11,12", values)
}

func TestMonthShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	timeUnit := month()
	values := []int{1, 12}

	err := timeUnitValidateMultipleInputs(timeUnit, values)

	assert.Equal(t, nil, err)
}

func TestMonthShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	timeUnit := month()
	value := 0

	err := timeUnitValidateInput(timeUnit, value)

	assert.Equal(t, "invalid month value 0 outside bounds 1 and 12", err.Error())
}

func TestMonthShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	timeUnit := month()
	value := 13

	err := timeUnitValidateInput(timeUnit, value)

	assert.Equal(t, "invalid month value 13 outside bounds 1 and 12", err.Error())
}

func TestDayOfWeekShouldReturnValues(t *testing.T) {
	timeUnit := dayOfWeek()

	allValues := timeUnitToAllValues(timeUnit)

	assert.Equal(t, inclusiveRange(0, 6), allValues)
}

func TestDayOfWeekShouldReturnTextualValuesConvertedToInts(t *testing.T) {
	timeUnit := dayOfWeek()
	inputValues := "MON,Tue,weD,tHu,fri,sat,sun"

	values := timeUnitInputToIntValues(timeUnit, inputValues)

	assert.Equal(t, "0,1,2,3,4,5,6", values)
}

func TestDayOfWeekShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	timeUnit := dayOfWeek()
	values := []int{0, 6}

	err := timeUnitValidateMultipleInputs(timeUnit, values)

	assert.Equal(t, nil, err)
}

func TestDayOfWeekShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	timeUnit := dayOfWeek()
	value := -1

	err := timeUnitValidateInput(timeUnit, value)

	assert.Equal(t, "invalid day of week value -1 outside bounds 0 and 6", err.Error())
}

func TestDayOfWeekShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	timeUnit := dayOfWeek()
	value := 7

	err := timeUnitValidateInput(timeUnit, value)

	assert.Equal(t, "invalid day of week value 7 outside bounds 0 and 6", err.Error())
}
