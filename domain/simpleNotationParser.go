package domain

import (
	"strconv"
)

type simpleNotationParser struct{}

func (p *simpleNotationParser) appliesTo(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func (p *simpleNotationParser) toValues(input string, timeUnit timeUnit) ([]int, error) {
	value, _ := strconv.Atoi(input)
	err := Validate(timeUnit, value)
	return []int{value}, err
}
