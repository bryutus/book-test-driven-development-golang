package chapter13

type sum struct {
	augend Money
	addend Money
}

func newSum(augend Money, addend Money) Expression {
	return sum{augend: augend, addend: addend}
}

func (s sum) getAugend() Money {
	return s.augend
}

func (s sum) getAddend() Money {
	return s.addend
}