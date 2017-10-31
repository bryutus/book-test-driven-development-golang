package chapter2

type Dollar struct {
	Amount int
}

func (d *Dollar) times(multiplier int) *Dollar {
	return newDollar(d.Amount * multiplier)
}

func newDollar(amount int) *Dollar {
	d := new(Dollar)
	d.Amount = amount
	return d
}
