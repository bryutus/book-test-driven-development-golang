package chapter6

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
				t.Errorf("Money.doller.times(): got %v want %v", got, want)
			}
		})
	}
}

func TestEquality(t *testing.T) {

	testCases := []struct {
		desc string
		in   int
		want bool
	}{
		{"5ドルと5ドルが同じ価値であること", 5, true},
		{"5ドルと6ドルが異なる価値であること", 6, false},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := newMoney(Dollar, 5).equals(newMoney(Dollar, test.in)); got != test.want {
				t.Errorf("Doller.equals(): got %v want %v", got, test.want)
			}
		})
	}

	testCases = []struct {
		desc string
		in   int
		want bool
	}{
		{"5CHFと5CHFが同じ価値であること", 5, true},
		{"5CHFと6CHFが異なる価値であること", 6, false},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := newMoney(Franc, 5).equals(newMoney(Franc, test.in)); got != test.want {
				t.Errorf("Franc.equals(): got %v want %v", got, test.want)
			}
		})
	}
}

func TestFrancMultiplication(t *testing.T) {

	testCases := []struct {
		desc string
		in   int
		want int
	}{
		{"5CHF*2=10となること", 2, 10},
		{"5CHF*3=15となること", 3, 15},
	}

	five := newMoney(Franc, 5)

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got, want := five.times(test.in), newMoney(Franc, test.want); got != want {
				t.Errorf("Money.franc.times(): got %v want %v", got, want)
			}
		})
	}
}
