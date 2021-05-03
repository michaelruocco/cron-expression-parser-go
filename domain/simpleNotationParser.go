package domain

import (
	"strconv"
)

type simpleNotationParser struct{}

func (p *simpleNotationParser) appliesTo(input string) bool {
	value, err := strconv.Atoi(input)
	return err == nil && value > -1
}

func (p *simpleNotationParser) toValues(input string, timeUnit timeUnit) ([]int, error) {
	value, parseErr := strconv.Atoi(input)
	if parseErr != nil {
		return nil, parseErr
	}
	err := Validate(timeUnit, value)
	return []int{value}, err
}
