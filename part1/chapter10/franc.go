package chapter10

type franc struct {
	money
}

func (f franc) times(multiplier int) Money {
	return newMoney(f.currency, f.amount*multiplier)
}
