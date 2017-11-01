package chapter4

type Dollar struct {
	Amount int
}

func (d *Dollar) times(multiplier int) *Dollar {
	return newDollar(d.Amount * multiplier)
}

func (d *Dollar) equals(dollar *Dollar) bool {
	return d.Amount == dollar.Amount
}

func newDollar(amount int) *Dollar {
	d := new(Dollar)
	d.Amount = amount
	return d
}
