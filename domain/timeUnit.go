package domain

import (
	"fmt"
)

type timeUnit struct {
	name       string
	lowerBound int
	upperBound int
}

func Minute() timeUnit {
	return timeUnit{"minute", 0, 59}
}

func Hour() timeUnit {
	return timeUnit{"hour", 0, 23}
}

func DayOfMonth() timeUnit {
	return timeUnit{"day of month", 1, 31}
}

func Month() timeUnit {
	return timeUnit{"month", 1, 12}
}

func DayOfWeek() timeUnit {
	return timeUnit{"day of week", 0, 6}
}

func AllValues(timeUnit timeUnit) []int {
	min := timeUnit.lowerBound
	values := make([]int, timeUnit.upperBound-min+1)
	for i := range values {
		values[i] = min + i
	}
	return values
}

func ValidateMultiple(timeUnit timeUnit, inputs []int) error {
	for _, input := range inputs {
		err := Validate(timeUnit, input)
		if err != nil {
			return err
		}
	}
	return nil
}

func Validate(timeUnit timeUnit, input int) error {
	if timeUnit.isOutOfBounds(input) {
		return fmt.Errorf("invalid %v value %d outside bounds %d and %d",
			timeUnit.name,
			input,
			timeUnit.lowerBound,
			timeUnit.upperBound)
	}
	return nil
}

func (timeUnit timeUnit) isOutOfBounds(input int) bool {
	return input < timeUnit.lowerBound || input > timeUnit.upperBound
}
