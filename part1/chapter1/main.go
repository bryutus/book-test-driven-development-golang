package main

type Dollar struct {
	Amount int
}

func (d *Dollar) times(multiplier int) {
	d.Amount = 10
}

func newDollar(amount int) *Dollar {
	d := new(Dollar)
	d.Amount = amount
	return d
}
