package chapter9

type franc struct {
	money
}

func (f franc) times(multiplier int) Money {
	return newMoney(Franc, f.amount*multiplier)
}

func (f franc) getCurrency() Currency {
	return "CHF"
}
