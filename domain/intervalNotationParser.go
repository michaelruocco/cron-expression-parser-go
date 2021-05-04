package domain

import (
	"fmt"
	"sort"
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
	starts, parseErr := p.toStarts(parts[0], timeUnit)
	if parseErr != nil {
		return nil, parseErr
	}
	interval, parseErr := strconv.Atoi(parts[1])
	if parseErr != nil {
		return nil, fmt.Errorf("invalid interval notation %v", input)
	}
	boundsErr := timeUnitValidateMultipleInputs(timeUnit, starts)
	if boundsErr != nil {
		return nil, boundsErr
	}
	return calculateIntervals(starts, timeUnit, interval), nil
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

func (p *intervalNotationParser) toStarts(input string, timeUnit timeUnit) ([]int, error) {
	if isWildcard(input) {
		return []int{timeUnit.lowerBound}, nil
	}
	if p.rangeParser.appliesTo(input) {
		return p.rangeParser.toValues(input, timeUnit)
	}
	return p.simpleParser.toValues(input, timeUnit)
}

func calculateIntervals(starts []int, timeUnit timeUnit, interval int) []int {
	intervals := []int{}
	for _, start := range starts {
		intervals = append(intervals, toInterval(start, timeUnit, interval)...)
	}
	intervals = unique(intervals)
	sort.Ints(intervals)
	return intervals
}

func toInterval(start int, timeUnit timeUnit, interval int) []int {
	intervals := []int{}
	for i := start; i <= timeUnit.upperBound; i = i + interval {
		intervals = append(intervals, i)
	}
	return intervals
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
