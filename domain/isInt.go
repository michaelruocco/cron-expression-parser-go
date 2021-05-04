package domain

import (
	"regexp"
)

var isInt = regexp.MustCompile(`^\d+$`).MatchString
