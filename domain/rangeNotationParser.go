package domain

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
		return nil, fmt.Errorf("invalid range notation %v", input)
	}
	start, parseErr := strconv.Atoi(parts[0])
	if parseErr != nil {
		return nil, fmt.Errorf("invalid range notation %v", input)
	}
	end, parseErr := strconv.Atoi(parts[1])
	if parseErr != nil {
		return nil, fmt.Errorf("invalid range notation %v", input)
	}
	boundsErr := timeUnit.validate([]int{start, end})
	if boundsErr != nil {
		return nil, boundsErr
	}
	return inclusiveRange(start, end), nil
}

func splitHyphen(input string) []string {
	return strings.Split(input, "-")
}
