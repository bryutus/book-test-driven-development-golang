package chapter6

type Franc struct {
	amount int
}

func (f Franc) times(multiplier int) Franc {
	return newFranc(f.amount * multiplier)
}

func (f Franc) equals(franc Franc) bool {
	return f.amount == franc.amount
}

func newFranc(amount int) Franc {
	return Franc{amount: amount}
}
