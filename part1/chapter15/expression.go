package chapter15

type Expression interface {
	getAugend() Expression
	getAddend() Expression
	reduce(Bank, Currency) Money
	plus(Expression) Expression
}
