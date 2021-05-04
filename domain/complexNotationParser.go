package domain

import (
	"strings"
)

func buildComplexNotationParser() notationParser {
	return &complexNotationParser{
		parsers: []notationParser{
			&simpleNotationParser{},
			&wildcardNotationParser{},
			&rangeNotationParser{},
			buildIntervalNotationParser(),
		},
	}
}

type complexNotationParser struct {
	parsers []notationParser
}

func (p *complexNotationParser) appliesTo(input string) bool {
	if appliesToSegment(p.parsers, input) {
		return true
	}
	segments := toSegments(input)
	for _, segment := range segments {
		if appliesToSegment(p.parsers, segment) {
			return true
		}
	}
	return false
}

func (p *complexNotationParser) toValues(input string, timeUnit timeUnit) ([]int, error) {
	if appliesToSegment(p.parsers, input) {
		return segmentToValues(p.parsers, input, timeUnit)
	}
	segments := toSegments(input)
	var allValues []int
	for _, segment := range segments {
		segmentValues, err := segmentToValues(p.parsers, segment, timeUnit)
		if err != nil {
			return nil, err
		}
		allValues = append(allValues, segmentValues...)
	}
	return allValues, nil
}

func toSegments(input string) []string {
	return strings.Split(input, ",")
}

func appliesToSegment(parsers []notationParser, segment string) bool {
	for _, parser := range parsers {
		if parser.appliesTo(segment) {
			return true
		}
	}
	return false
}

func segmentToValues(parsers []notationParser, segment string, timeUnit timeUnit) ([]int, error) {
	var values []int
	for _, parser := range parsers {
		if parser.appliesTo(segment) {
			segmentValues, err := parser.toValues(segment, timeUnit)
			if err != nil {
				return nil, err
			}
			values = append(values, segmentValues...)
		}
	}
	return values, nil
}
