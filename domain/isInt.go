package domain

import (
	"regexp"
)

var IsInt = regexp.MustCompile(`^\d+$`).MatchString
