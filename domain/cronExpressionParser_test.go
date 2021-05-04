package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	args := []string{"1", "1", "1", "JAN-DEC", "MON-SUN", "my-command-5"}

	output, _ := ParseCronExpression(args)

	assert.Equal(t, "minute        1\n"+
		"hour          1\n"+
		"day of month  1\n"+
		"month         1 2 3 4 5 6 7 8 9 10 11 12\n"+
		"day of week   0 1 2 3 4 5 6\n"+
		"command       my-command-5", output)
}
