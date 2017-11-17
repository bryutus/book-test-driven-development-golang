package chapter15

type Pair interface{}

type pair struct {
	from Currency
	to   Currency
}

func newPair(from Currency, to Currency) Pair {
	return pair{from: from, to: to}
}
