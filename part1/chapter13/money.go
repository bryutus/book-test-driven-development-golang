package chapter13

type Currency string

const (
	Dollar = "USD"
	Franc  = "CHF"
)

type Money interface {
	times(int) Money
	equals(Money) bool
	plus(Money) Expression
	getAmount() int
	getCurrency() Currency
}

type money struct {
	amount   int
	currency Currency
}

func newMoney(currency Currency, amount int) Money {
	return money{amount: amount, currency: currency}
}

func (m money) times(multiplier int) Money {
	return newMoney(m.currency, m.amount*multiplier)
}

func (m money) equals(money Money) bool {
	return m.amount == money.getAmount() && m.currency == money.getCurrency()
}

func (m money) plus(addend Money) Expression {
	return newSum(m, addend)
}

func (m money) getAmount() int {
	return m.amount
}

func (m money) getCurrency() Currency {
	return m.currency
}
