package chapter14

type Bank interface {
	reduce(Expression, Currency) Money
	addRate(Currency, Currency, int)
	rate(Currency, Currency) int
}

type bank struct {
	rates map[Pair]int
}

func newBank() Bank {
	return &bank{}
}

func (b *bank) reduce(source Expression, to Currency) Money {
	return source.reduce(b, to)
}

func (b *bank) addRate(from Currency, to Currency, rate int) {
	p := newPair(from, to)
	b.rates = map[Pair]int{p: rate}
}

func (b *bank) rate(from Currency, to Currency) int {
	if from == to {
		return 1
	}
	return b.rates[newPair(from, to)]
}
