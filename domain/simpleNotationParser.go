package domain

import (
	"fmt"
	"strconv"
	"strings"
)

type simpleNotationParser struct{}

func (p *simpleNotationParser) appliesTo(input string) bool {
	return isInts(splitComma(input))
}

func (p *simpleNotationParser) toValues(input string, timeUnit timeUnit) ([]int, error) {
	values, parseErr := toInts(input)
	if parseErr != nil {
		return nil, fmt.Errorf("invalid notation %v", input)
	}
	err := timeUnit.validate(values)
	return values, err
}

func toInts(input string) ([]int, error) {
	segments := splitComma(input)
	values := make([]int, len(segments))
	for i, segment := range segments {
		value, parseErr := strconv.Atoi(segment)
		if parseErr != nil {
			return nil, fmt.Errorf("invalid notation %v", input)
		}
		values[i] = value
	}
	return values, nil
}

func splitComma(input string) []string {
	return strings.Split(input, ",")
}
