package chapter14

type Bank interface {
	reduce(Expression, Currency) Money
	addRate(Currency, Currency, int)
	rate(Currency, Currency) int
}

type bank struct {
}

func newBank() Bank {
	return bank{}
}

func (b bank) reduce(source Expression, to Currency) Money {
	return source.reduce(b, to)
}

func (b bank) addRate(from Currency, to Currency, rate int) {}

func (b bank) rate(from Currency, to Currency) (rate int) {
	rate = 1
	if from == Franc && to == Dollar {
		rate = 2
	}
	return
}
