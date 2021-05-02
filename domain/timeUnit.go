package domain

type timeUnit struct {
	lowerBound int
	upperBound int
}

func Minute() timeUnit {
	return timeUnit{lowerBound: 0, upperBound: 59}
}

func Hour() timeUnit {
	return timeUnit{lowerBound: 0, upperBound: 23}
}

func DayOfMonth() timeUnit {
	return timeUnit{lowerBound: 1, upperBound: 31}
}

func Month() timeUnit {
	return timeUnit{lowerBound: 1, upperBound: 12}
}

func DayOfWeek() timeUnit {
	return timeUnit{lowerBound: 0, upperBound: 6}
}

func AllValues(timeUnit timeUnit) []int {
	min := timeUnit.lowerBound
	values := make([]int, timeUnit.upperBound-min+1)
	for i := range values {
		values[i] = min + i
	}
	return values
}
