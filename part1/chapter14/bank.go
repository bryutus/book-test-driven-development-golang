package chapter14

type Bank interface {
	reduce(Expression, Currency) Money
}

type bank struct {
}

func newBank() Bank {
	return bank{}
}

func (b bank) reduce(source Expression, to Currency) Money {
	return source.reduce(to)
}
