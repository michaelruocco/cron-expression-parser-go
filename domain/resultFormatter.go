package domain

import (
	"fmt"
	"sort"
	"strings"
)

func format(result cronResult) string {
	return strings.Join(toLines(result), "\n")
}

func toLines(result cronResult) []string {
	return []string{
		formatMinutes(result),
		formatHours(result),
		formatDaysOfMonth(result),
		formatMonths(result),
		formatDaysOfWeek(result),
		formatCommand(result),
	}
}

func formatMinutes(result cronResult) string {
	return formatValues("minute", toString(result.Minutes))
}

func formatHours(result cronResult) string {
	return formatValues("hour", toString(result.Hours))
}

func formatDaysOfMonth(result cronResult) string {
	return formatValues("day of month", toString(result.DaysOfMonth))
}

func formatMonths(result cronResult) string {
	return formatValues("month", toString(result.Months))
}

func formatDaysOfWeek(result cronResult) string {
	return formatValues("day of week", toString(result.DaysOfWeek))
}

func formatCommand(result cronResult) string {
	return formatValues("command", result.Command)
}

func formatValues(name string, values string) string {
	return fmt.Sprintf("%-14v%v", name, values)
}

func toString(values []int) string {
	sortedValues := make([]int, len(values))
	copy(sortedValues, values)
	sort.Ints(sortedValues)
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sortedValues)), " "), "[]")
}
