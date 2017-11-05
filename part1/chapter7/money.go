package chapter7

type Currency int

const (
	Dollar = iota
	Franc
)

type Money interface {
	times(int) Money
	equals(Money) bool
	getAmount() int
	getCurrency() Currency
}

type money struct {
	amount   int
	currency Currency
}

func newMoney(currency Currency, amount int) Money {
	m := money{amount: amount, currency: currency}

	if currency == Franc {
		return franc{m}
	}

	return dollar{m}
}

func (m money) equals(money Money) bool {
	return m.amount == money.getAmount() && m.currency == money.getCurrency()
}

func (m money) getAmount() int {
	return m.amount
}

func (m money) getCurrency() Currency {
	return m.currency
}
