package chapter6

type Dollar struct {
	Money
}

func (d Dollar) times(multiplier int) Dollar {
	return newDollar(d.amount * multiplier)
}

func (d Dollar) equals(dollar Dollar) bool {
	return d.amount == dollar.amount
}

func newDollar(amount int) Dollar {
	return Dollar{Money{amount: amount}}
}