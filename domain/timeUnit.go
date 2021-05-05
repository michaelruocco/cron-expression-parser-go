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

func minute() timeUnit {
	return timeUnit{"minute", 0, 59}
}

func hour() timeUnit {
	return timeUnit{"hour", 0, 23}
}

func dayOfMonth() timeUnit {
	return timeUnit{"day of month", 1, 31}
}

func month() timeUnit {
	return timeUnit{"month", 1, 12}
}

func dayOfWeek() timeUnit {
	return timeUnit{"day of week", 0, 6}
}

func (u *timeUnit) allValues() []int {
	min := u.lowerBound
	values := make([]int, u.upperBound-min+1)
	for i := range values {
		values[i] = min + i
	}
	return values
}

func (u *timeUnit) toIntValues(inputValues string) string {
	switch u.name {
	case "day of week":
		return toDaysOfWeek(inputValues)
	case "month":
		return toMonths(inputValues)
	default:
		return inputValues
	}
}

func (u *timeUnit) validate(inputs []int) error {
	for _, input := range inputs {
		if u.isOutOfBounds(input) {
			return fmt.Errorf("invalid %v value %d outside bounds %d and %d",
				u.name,
				input,
				u.lowerBound,
				u.upperBound)
		}
	}
	return nil
}

func (u *timeUnit) isOutOfBounds(input int) bool {
	return input < u.lowerBound || input > u.upperBound
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
