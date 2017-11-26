package chapter16

type Expression interface {
	getAugend() Expression
	getAddend() Expression
	reduce(Bank, Currency) Money
	plus(Expression) Expression
	times(int) Expression
}
