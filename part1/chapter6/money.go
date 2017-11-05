package chapter6

type Currency int

const (
	Dollar = iota
	Franc
)

type Money interface {
	times(int) Money
	equals(Money) bool
	getAmount() int
}

type money struct {
	amount int
}

func newMoney(currency Currency, amount int) Money {
	m := money{amount: amount}

	if currency == Franc {
		return franc{m}
	}

	return dollar{m}
}

func (m money) equals(money Money) bool {
	return m.amount == money.getAmount()
}

func (m money) getAmount() int {
	return m.amount
}
