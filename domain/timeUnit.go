package domain

import (
	"fmt"
	"strings"
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

func ToIntValues(timeUnit timeUnit, inputValues string) string {
	switch timeUnit.name {
	case "day of week":
		return toDaysOfWeek(inputValues)
	case "month":
		return toMonths(inputValues)
	default:
		return inputValues
	}
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
	if isOutOfBounds(timeUnit, input) {
		return fmt.Errorf("invalid %v value %d outside bounds %d and %d",
			timeUnit.name,
			input,
			timeUnit.lowerBound,
			timeUnit.upperBound)
	}
	return nil
}

func toDaysOfWeek(input string) string {
	replacer := strings.NewReplacer(
		"mon", "0",
		"tue", "1",
		"wed", "2",
		"thu", "3",
		"fri", "4",
		"sat", "5",
		"sun", "6",
	)
	return replacer.Replace(strings.ToLower(input))
}

func toMonths(input string) string {
	replacer := strings.NewReplacer(
		"jan", "1",
		"feb", "2",
		"mar", "3",
		"apr", "4",
		"may", "5",
		"jun", "6",
		"jul", "7",
		"aug", "8",
		"sep", "9",
		"oct", "10",
		"nov", "11",
		"dec", "12",
	)
	return replacer.Replace(strings.ToLower(input))
}

func isOutOfBounds(timeUnit timeUnit, input int) bool {
	return input < timeUnit.lowerBound || input > timeUnit.upperBound
}
