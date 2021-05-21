package parser

import (
	"fmt"
	"strconv"
	"strings"
)

func buildIntervalNotationParser() notationParser {
	return &intervalNotationParser{
		simpleParser: &simpleNotationParser{},
		rangeParser:  &rangeNotationParser{},
	}
}

type intervalNotationParser struct {
	simpleParser notationParser
	rangeParser  notationParser
}

func (p *intervalNotationParser) appliesTo(input string) bool {
	parts := splitSlash(input)
	if len(parts) == 2 {
		return p.isIntOrWildcardOrRange(parts[0]) && isInt(parts[1])
	}
	return false
}

func (p *intervalNotationParser) toValues(input string, timeUnit timeUnit) ([]int, error) {
	parts := splitSlash(input)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid interval notation %v", input)
	}
	start, parseErr := p.toStart(parts[0], timeUnit)
	if parseErr != nil {
		return nil, parseErr
	}
	end, parseErr := p.toEnd(parts[0], timeUnit)
	interval, parseErr := strconv.Atoi(parts[1])
	startErr := timeUnit.validate([]int{start})
	endErr := timeUnit.validate([]int{end})
	if parseErr != nil || startErr != nil || endErr != nil {
		return nil, fmt.Errorf("invalid interval notation %v", input)
	}

	return calculateIntervals(start, end, interval), nil
}

func splitSlash(input string) []string {
	return strings.Split(input, "/")
}

func (p *intervalNotationParser) isIntOrWildcardOrRange(input string) bool {
	return isWildcard(input) ||
		p.rangeParser.appliesTo(input) ||
		p.simpleParser.appliesTo(input)
}

func isWildcard(input string) bool {
	return input == "*"
}

func (p *intervalNotationParser) toStart(input string, timeUnit timeUnit) (int, error) {
	if isWildcard(input) {
		return timeUnit.lowerBound, nil
	}
	if p.rangeParser.appliesTo(input) {
		return extractFirst(input, timeUnit, p.rangeParser)
	}
	return extractFirst(input, timeUnit, p.simpleParser)
}

func extractFirst(input string, timeUnit timeUnit, parser notationParser) (int, error) {
	values, err := parser.toValues(input, timeUnit)
	if err != nil {
		return -1, err
	}
	return values[0], nil
}

func (p *intervalNotationParser) toEnd(input string, timeUnit timeUnit) (int, error) {
	if p.rangeParser.appliesTo(input) {
		values, _ := p.rangeParser.toValues(input, timeUnit)
		return values[len(values)-1], nil
	}
	return timeUnit.upperBound, nil
}

func calculateIntervals(start int, end int, interval int) []int {
	intervals := []int{}
	for i := start; i <= end; i = i + interval {
		intervals = append(intervals, i)
	}
	return unique(intervals)
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
