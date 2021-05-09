package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const command = "my-command"

func TestShouldHandleSimpleNotationForAllTimeUnits(t *testing.T) {
	args := []string{"30", "12", "15", "6", "3", "my-command-1"}

	output, _ := ParseCronExpression(args)

	assert.Equal(t, "minute        30\n"+
		"hour          12\n"+
		"day of month  15\n"+
		"month         6\n"+
		"day of week   3\n"+
		"command       my-command-1", output)
}

func TestShouldHandleWildcardNotationForAllTimeUnits(t *testing.T) {
	args := []string{"*", "*", "*", "*", "*", "my-command-2"}

	output, _ := ParseCronExpression(args)

	assert.Equal(t, "minute        0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59\n"+
		"hour          0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23\n"+
		"day of month  1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31\n"+
		"month         1 2 3 4 5 6 7 8 9 10 11 12\n"+
		"day of week   0 1 2 3 4 5 6\n"+
		"command       my-command-2", output)
}

func TestShouldHandleRangeNotationForAllTimeUnits(t *testing.T) {
	args := []string{"30-32", "12-14", "15-17", "6-8", "3-5", "my-command-3"}

	output, _ := ParseCronExpression(args)

	assert.Equal(t, "minute        30 31 32\n"+
		"hour          12 13 14\n"+
		"day of month  15 16 17\n"+
		"month         6 7 8\n"+
		"day of week   3 4 5\n"+
		"command       my-command-3", output)
}

func TestShouldHandleCommaNotationForAllTimeUnits(t *testing.T) {
	args := []string{"30,32", "12,14", "15,17", "6,8", "3,5", "my-command-4"}

	output, _ := ParseCronExpression(args)

	assert.Equal(t, "minute        30 32\n"+
		"hour          12 14\n"+
		"day of month  15 17\n"+
		"month         6 8\n"+
		"day of week   3 5\n"+
		"command       my-command-4", output)
}

func TestShouldHandleTextualInputsForMonthAndDayOfWeek(t *testing.T) {
	args := []string{"1", "1", "1", "JAN-FEB", "SAT-SUN", "my-command-5"}

	output, _ := ParseCronExpression(args)

	assert.Equal(t, "minute        1\n"+
		"hour          1\n"+
		"day of month  1\n"+
		"month         1 2\n"+
		"day of week   5 6\n"+
		"command       my-command-5", output)
}

func TestShouldHandleComplexNotation(t *testing.T) {
	args := []string{"1,4-8,*/15", "2", "2", "MAR-MAY", "MON-TUE", "my-command-6"}

	output, _ := ParseCronExpression(args)

	assert.Equal(t, "minute        0 1 4 5 6 7 8 15 30 45\n"+
		"hour          2\n"+
		"day of month  2\n"+
		"month         3 4 5\n"+
		"day of week   0 1\n"+
		"command       my-command-6", output)
}

func TestShouldHandleComplexIntervalStartNotation(t *testing.T) {
	args := []string{"3,45/15", "3", "3", "JUN-JUL", "WED-THU", "my-command-7"}

	output, _ := ParseCronExpression(args)

	assert.Equal(t, "minute        3 18 33 45 48\n"+
		"hour          3\n"+
		"day of month  3\n"+
		"month         6 7\n"+
		"day of week   2 3\n"+
		"command       my-command-7", output)
}

func TestShouldHandleArgumentOption(t *testing.T) {
	args := []string{"-arguments", "4", "4", "4", "1", "1", "my-command-6"}

	output, _ := ParseCronExpression(args)

	assert.Equal(t, "minute        4\n"+
		"hour          4\n"+
		"day of month  4\n"+
		"month         1\n"+
		"day of week   1\n"+
		"command       my-command-6", output)
}

func TestShouldReturnErrorIfRequiredArgumentsNotProvided(t *testing.T) {
	args := []string{}

	_, err := ParseCronExpression(args)

	assert.Equal(t, "usage: please provide a valid cron expression", err.Error())
}

func TestShouldNotAllowMinuteValueOutsideBounds(t *testing.T) {
	args := []string{"61", "1", "1", "1", "1", command}

	_, err := ParseCronExpression(args)

	assert.Equal(t, "invalid minute value 61 outside bounds 0 and 59", err.Error())
}

func TestShouldNotAllowHourValueOutsideBounds(t *testing.T) {
	args := []string{"1", "24", "1", "1", "1", command}

	_, err := ParseCronExpression(args)

	assert.Equal(t, "invalid hour value 24 outside bounds 0 and 23", err.Error())
}

func TestShouldNotAllowDayOfMonthValueOutsideBounds(t *testing.T) {
	args := []string{"1", "1", "32", "1", "1", command}

	_, err := ParseCronExpression(args)

	assert.Equal(t, "invalid day of month value 32 outside bounds 1 and 31", err.Error())
}

func TestShouldNotAllowMonthValueOutsideBounds(t *testing.T) {
	args := []string{"1", "1", "1", "13", "1", command}

	_, err := ParseCronExpression(args)

	assert.Equal(t, "invalid month value 13 outside bounds 1 and 12", err.Error())
}

func TestShouldNotAllowDayOfWeekValueOutsideBounds(t *testing.T) {
	args := []string{"1", "1", "1", "1", "7", command}

	_, err := ParseCronExpression(args)

	assert.Equal(t, "invalid day of week value 7 outside bounds 0 and 6", err.Error())
}

func TestShouldNotAllowNotIntegerValuesForMinute(t *testing.T) {
	args := []string{"1.1", "1", "1", "1", "7", command}

	_, err := ParseCronExpression(args)

	assert.Equal(t, "notation parser not found for value 1.1", err.Error())
}

func TestShouldNotAllowNotIntegerValuesForHour(t *testing.T) {
	args := []string{"1", "1.2", "1", "1", "7", command}

	_, err := ParseCronExpression(args)

	assert.Equal(t, "notation parser not found for value 1.2", err.Error())
}

func TestShouldNotAllowNotIntegerValuesForDayOfMonth(t *testing.T) {
	args := []string{"1", "1", "1.3", "1", "7", command}

	_, err := ParseCronExpression(args)

	assert.Equal(t, "notation parser not found for value 1.3", err.Error())
}

func TestShouldNotAllowNotIntegerValuesForMonth(t *testing.T) {
	args := []string{"1", "1", "1", "1.4", "7", command}

	_, err := ParseCronExpression(args)

	assert.Equal(t, "notation parser not found for value 1.4", err.Error())
}

func TestShouldNotAllowNotIntegerValuesForDayOfWeek(t *testing.T) {
	args := []string{"1", "1", "1", "1", "1.5", command}

	_, err := ParseCronExpression(args)

	assert.Equal(t, "notation parser not found for value 1.5", err.Error())
}
