package chapter1

type Dollar struct {
	Amount int
}

func (d *Dollar) times(multiplier int) {
	d.Amount *= multiplier
}

func newDollar(amount int) *Dollar {
	d := new(Dollar)
	d.Amount = amount
	return d
}
