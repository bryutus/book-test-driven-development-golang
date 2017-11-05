package chapter6

type franc struct {
	money
}

func (f franc) times(multiplier int) Money {
	return newMoney(Franc, f.amount*multiplier)
}
