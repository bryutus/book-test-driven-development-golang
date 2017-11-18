package chapter15

type Currency string

const (
	Dollar = "USD"
	Franc  = "CHF"
)

type Money interface {
	times(int) Expression
	equals(Money) bool
	getAmount() int
	getCurrency() Currency
	Expression
}

type money struct {
	amount   int
	currency Currency
}

func newMoney(currency Currency, amount int) Money {
	return money{amount: amount, currency: currency}
}

func (m money) times(multiplier int) Expression {
	return newMoney(m.currency, m.amount*multiplier)
}

func (m money) equals(money Money) bool {
	return m.amount == money.getAmount() && m.currency == money.getCurrency()
}

func (m money) plus(addend Expression) Expression {
	return newSum(m, addend)
}

func (m money) getAmount() int {
	return m.amount
}

func (m money) getCurrency() Currency {
	return m.currency
}

func (m money) getAugend() Expression {
	return m
}

func (m money) getAddend() Expression {
	return m
}

func (m money) reduce(bank Bank, to Currency) Money {
	rate := bank.rate(m.currency, to)
	return newMoney(to, m.amount/rate)
}
