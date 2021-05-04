package domain

import (
	"strconv"
)

type simpleNotationParser struct{}

func (p *simpleNotationParser) appliesTo(input string) bool {
	return isInt(input)
}

func (p *simpleNotationParser) toValues(input string, timeUnit timeUnit) ([]int, error) {
	value, parseErr := strconv.Atoi(input)
	if parseErr != nil {
		return nil, parseErr
	}
	err := timeUnitValidateInput(timeUnit, value)
	return []int{value}, err
}
