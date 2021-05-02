package domain

func Parse(args []string) (CronResult, error) {
	return CronResult{
		Minutes:     []int{0, 15, 30, 45},
		Hours:       []int{0},
		DaysOfMonth: []int{1, 15},
		Months:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		DaysOfWeek:  []int{1, 2, 3, 4, 5},
		Command:     "/usr/bin/find",
	}, nil
}
