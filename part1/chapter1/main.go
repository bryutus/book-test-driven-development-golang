package main

type Dollar struct {
	Amount int
}

func (d *Dollar) times(multiplier int) {
	d.Amount = 10
}
