package chapter15

type Expression interface {
	getAugend() Money
	getAddend() Money
	reduce(Bank, Currency) Money
}
