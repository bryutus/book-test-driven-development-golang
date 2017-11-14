package chapter14

type Bank interface {
	reduce(Expression, Currency) Money
	addRate(Currency, Currency, int)
}

type bank struct {
}

func newBank() Bank {
	return bank{}
}

func (b bank) reduce(source Expression, to Currency) Money {
	return source.reduce(to)
}

func (b bank) addRate(from Currency, to Currency, rate int) {}
