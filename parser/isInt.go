package parser

import (
	"regexp"
)

var isInt = regexp.MustCompile(`^\d+$`).MatchString

func isInts(input []string) bool {
	for _, value := range input {
		if !isInt(value) {
			return false
		}
	}
	return true
}
