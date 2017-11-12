package chapter13

type Bank interface {
	reduce(Expression, Currency) Money
}

type bank struct {
}

func newBank() Bank {
	return bank{}
}

func (b bank) reduce(source Expression, to Currency) Money {
	amount := source.getAugend().getAmount() + source.getAddend().getAmount()
	return newMoney(to, amount)
}
