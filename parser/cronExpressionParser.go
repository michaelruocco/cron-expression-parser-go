package parser

import (
	"fmt"
)

//Takes a set of arguments which comprise the cron expression split by spaces.
//Returns a cron result which contains integer values for the months, days and hours
//along with the command to be executes, or will return an error if the input arguments
//cannot be successfully parsed
func ParseCronExpression(args []string) (string, error) {
	santizedArgs, err := sanitize(args)
	if err != nil {
		return "", err
	}

	result, err := parseExpression(santizedArgs)
	if err != nil {
		return "", err
	}

	return format(result), nil
}

func parseExpression(args []string) (cronResult, error) {
	notationParser := buildComplexNotationParser()

	minutes, err := parse(args[0], minute(), notationParser)
	if err != nil {
		return cronResult{}, err
	}

	hours, err := parse(args[1], hour(), notationParser)
	if err != nil {
		return cronResult{}, err
	}

	daysOfMonth, err := parse(args[2], dayOfMonth(), notationParser)
	if err != nil {
		return cronResult{}, err
	}

	months, err := parse(args[3], month(), notationParser)
	if err != nil {
		return cronResult{}, err
	}

	daysOfWeek, err := parse(args[4], dayOfWeek(), notationParser)
	if err != nil {
		return cronResult{}, err
	}

	return cronResult{
		Minutes:     minutes,
		Hours:       hours,
		DaysOfMonth: daysOfMonth,
		Months:      months,
		DaysOfWeek:  daysOfWeek,
		Command:     args[5],
	}, nil
}

func parse(rawValues string, timeUnit timeUnit, notationParser notationParser) ([]int, error) {
	intValues := timeUnit.toIntValues(rawValues)
	if notationParser.appliesTo(intValues) {
		values, err := notationParser.toValues(intValues, timeUnit)
		if err != nil {
			return nil, err
		}
		return values, nil
	}
	return nil, fmt.Errorf("notation parser not found for value %v", rawValues)
}
