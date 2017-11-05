package chapter6

type dollar struct {
	money
}

func (d dollar) times(multiplier int) Money {
	return newMoney(Dollar, d.amount*multiplier)
}
