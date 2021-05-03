package domain

import (
	"fmt"
)

func Parse(args []string) (CronResult, error) {
	notationParser := BuildNotationParser()

	minutes, err := parse(args[0], Minute(), notationParser)
	if err != nil {
		return CronResult{}, err
	}

	hours, err := parse(args[1], Hour(), notationParser)
	if err != nil {
		return CronResult{}, err
	}

	daysOfMonth, err := parse(args[2], DayOfMonth(), notationParser)
	if err != nil {
		return CronResult{}, err
	}

	months, err := parse(args[3], Month(), notationParser)
	if err != nil {
		return CronResult{}, err
	}

	daysOfWeek, err := parse(args[4], DayOfWeek(), notationParser)
	if err != nil {
		return CronResult{}, err
	}

	return CronResult{
		Minutes:     minutes,
		Hours:       hours,
		DaysOfMonth: daysOfMonth,
		Months:      months,
		DaysOfWeek:  daysOfWeek,
		Command:     "/usr/bin/find",
	}, nil
}

func parse(rawValues string, timeUnit timeUnit, notationParser notationParser) ([]int, error) {
	intValues := ToIntValues(timeUnit, rawValues)
	fmt.Println("INT VALUES", intValues)
	if notationParser.appliesTo(intValues) {
		values, err := notationParser.toValues(intValues, timeUnit)
		if err != nil {
			return nil, err
		}
		return values, nil
	}
	return nil, fmt.Errorf("notation parser not found for value %v", rawValues)
}
