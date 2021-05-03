package domain

import (
	"strconv"
)

type simpleNotationParser struct{}

func (p *simpleNotationParser) appliesTo(input string) bool {
	return IsInt(input)
}

func (p *simpleNotationParser) toValues(input string, timeUnit timeUnit) ([]int, error) {
	value, parseErr := strconv.Atoi(input)
	if parseErr != nil {
		return nil, parseErr
	}
	err := Validate(timeUnit, value)
	return []int{value}, err
}
