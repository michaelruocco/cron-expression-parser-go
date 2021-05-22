package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonthShouldReturnAllValues(t *testing.T) {
	month := month()

	allValues := month.allValues()

	assert.Equal(t, inclusiveRange(1, 12), allValues)
}

func TestMonthShouldReturnTextualValuesConvertedToInts(t *testing.T) {
	month := month()
	inputValues := "JAN,Feb,maR,aPr,may,jun,jul,aug,sep,oct,nov,dec"

	values := month.toIntValues(inputValues)

	assert.Equal(t, "1,2,3,4,5,6,7,8,9,10,11,12", values)
}

func TestMonthShouldNotReturnErrorIfValuesAreWithinBounds(t *testing.T) {
	month := month()
	values := []int{1, 12}

	err := month.validate(values...)

	assert.Equal(t, nil, err)
}

func TestMonthShouldReturnErrorIfValueIsLessThanLowerBound(t *testing.T) {
	month := month()
	values := 0

	err := month.validate(values)

	assert.Equal(t, "invalid month value 0 outside bounds 1 and 12", err.Error())
}

func TestMonthShouldReturnErrorIfValueIsGreaterThanUpperBound(t *testing.T) {
	month := month()
	values := 13

	err := month.validate(values)

	assert.Equal(t, "invalid month value 13 outside bounds 1 and 12", err.Error())
}
