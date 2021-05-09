package parser

type notationParser interface {
	appliesTo(input string) bool

	toValues(input string, timeUnit timeUnit) ([]int, error)
}
