package chapter10

type dollar struct {
	money
}

func (d dollar) times(multiplier int) Money {
	return newMoney(d.currency, d.amount*multiplier)
}
