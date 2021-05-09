package parser

import (
	"errors"
)

const requiredArguments = 6

func sanitize(args []string) ([]string, error) {
	if len(args) < requiredArguments {
		return nil, errors.New("usage: please provide a valid cron expression")
	}
	return toRequiredArgs(args), nil
}

func toRequiredArgs(args []string) []string {
	size := len(args)
	return args[size-requiredArguments : size]
}
