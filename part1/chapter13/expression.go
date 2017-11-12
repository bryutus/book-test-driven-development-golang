package chapter13

type Expression interface {
	getAugend() Money
	getAddend() Money
}
