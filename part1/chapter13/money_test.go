package chapter13

import (
	"testing"
)

func TestMultiplication(t *testing.T) {

	testCases := []struct {
		desc string
		in   int
		want int
	}{
		{"$5*2=10となること", 2, 10},
		{"$5*3=15となること", 3, 15},
	}

	five := newMoney(Dollar, 5)

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got, want := five.times(test.in), newMoney(Dollar, test.want); got != want {
				t.Errorf("Money.times(): got %v want %v", got, want)
			}
		})
	}
}

func TestEquality(t *testing.T) {

	testCases := []struct {
		desc         string
		currencySrc  Currency
		currencyDest Currency
		in           int
		want         bool
	}{
		{"5ドルと5ドルが同じ価値であること", Dollar, Dollar, 5, true},
		{"5ドルと6ドルが異なる価値であること", Dollar, Dollar, 6, false},
		{"5ドルと5CHFが異なる価値であること", Dollar, Franc, 5, false},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := newMoney(test.currencySrc, 5).equals(newMoney(test.currencyDest, test.in)); got != test.want {
				t.Errorf("Money.equals(): got %v want %v", got, test.want)
			}
		})
	}
}

func TestCurrency(t *testing.T) {

	testCases := []struct {
		desc string
		in   Currency
		want Currency
	}{
		{"Dollarの通貨がUSDとなること", Dollar, "USD"},
		{"Francの通貨がCHFとなること", Franc, "CHF"},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := newMoney(test.in, 1).getCurrency(); got != test.want {
				t.Errorf("Money.getCurrency(): got %v want %v", got, test.want)
			}
		})
	}
}

func TestSimpleAddition(t *testing.T) {

	testCases := []struct {
		desc string
		a    int
		want int
	}{
		{"5ドル足す5ドルが10ドルとなること", 5, 10},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			five := newMoney(Dollar, test.a)
			sum := five.plus(five)
			bank := newBank()
			reduced := bank.reduce(sum, Dollar)
			if got, want := reduced, newMoney(Dollar, test.want); got != want {
				t.Errorf("Money.plus(): got %v want %v", got, want)
			}
		})
	}
}

func TestPlusResultSmu(t *testing.T) {
	five := newMoney(Dollar, 5)
	sum := five.plus(five)
	bank := newBank()
	if got, want := sum.augend, five; got != want {
		t.Errorf("Money.plus(): got %v want %v", got, want)
	}
	if got, want := sum.addend, five; got != want {
		t.Errorf("Money.plus(): got %v want %v", got, want)
	}

}
