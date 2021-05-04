package domain

type cronResult struct {
	Minutes     []int
	Hours       []int
	DaysOfMonth []int
	Months      []int
	DaysOfWeek  []int
	Command     string
}
