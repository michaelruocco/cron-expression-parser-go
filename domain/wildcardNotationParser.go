package domain

type wildcardNotationParser struct{}

func (p *wildcardNotationParser) appliesTo(input string) bool {
	return input == "*"
}

func (p *wildcardNotationParser) toValues(input string, timeUnit timeUnit) ([]int, error) {
	return timeUnitToAllValues(timeUnit), nil
}
