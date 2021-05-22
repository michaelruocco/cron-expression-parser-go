package parser

import (
	"fmt"
	"strconv"
	"strings"
)

type rangeNotationParser struct{}

func (p *rangeNotationParser) appliesTo(input string) bool {
	parts := splitHyphen(input)
	if len(parts) == 2 {
		return isInt(parts[0]) && isInt(parts[1])
	}
	return false
}

func (p *rangeNotationParser) toValues(input string, timeUnit timeUnit) ([]int, error) {
	parts := splitHyphen(input)
	if len(parts) != 2 {
		return nil, toInvalidRangeNotationError(input)
	}
	start, parseErr := strconv.Atoi(parts[0])
	if parseErr != nil {
		return nil, toInvalidRangeNotationError(input)
	}
	end, parseErr := strconv.Atoi(parts[1])
	if parseErr != nil {
		return nil, toInvalidRangeNotationError(input)
	}
	boundsErr := timeUnit.validate(start, end)
	if boundsErr != nil {
		return nil, boundsErr
	}
	return inclusiveRange(start, end), nil
}

func toInvalidRangeNotationError(input string) error {
	return fmt.Errorf("invalid range notation %v", input)
}

func splitHyphen(input string) []string {
	return strings.Split(input, "-")
}
