package chapter9

type dollar struct {
	money
}

func (d dollar) times(multiplier int) Money {
	return newMoney(Dollar, d.amount*multiplier)
}

func (d dollar) getCurrency() Currency {
	return "USD"
}
