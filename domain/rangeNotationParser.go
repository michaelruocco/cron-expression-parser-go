package domain

import (
	"fmt"
	"strconv"
	"strings"
)

type rangeNotationParser struct{}

func (p *rangeNotationParser) appliesTo(input string) bool {
	parts := split(input)
	if len(parts) == 2 {
		return IsInt(parts[0]) && IsInt(parts[1])
	}
	return false
}

func (p *rangeNotationParser) toValues(input string, timeUnit timeUnit) ([]int, error) {
	parts := split(input)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid range notation %v", input)
	}
	start, parseErr := strconv.Atoi(parts[0])
	if parseErr != nil {
		return nil, parseErr
	}
	end, parseErr := strconv.Atoi(parts[1])
	if parseErr != nil {
		return nil, parseErr
	}
	boundsErr := ValidateMultiple(timeUnit, []int{start, end})
	if boundsErr != nil {
		return nil, boundsErr
	}
	return InclusiveRange(start, end), nil
}

func split(input string) []string {
	return strings.Split(input, "-")
}