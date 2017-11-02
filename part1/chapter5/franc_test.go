package chapter5

import (
	"testing"
)

func TestFrancMultiplication(t *testing.T) {

	testCases := []struct {
		desc string
		in   int
		want int
	}{
		{"5CHF*2=10となること", 2, 10},
		{"5CHF*3=15となること", 3, 15},
	}

	five := newFranc(5)

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got, want := five.times(test.in), newFranc(test.want); got != want {
				t.Errorf("Franc.times(): got %v want %v", got, want)
			}
		})
	}
}

func TestFrancEquality(t *testing.T) {

	testCases := []struct {
		desc string
		in   int
		want bool
	}{
		{"5CHFと5CHFが同じ価値であること", 5, true},
		{"5CHFと6CHFが異なる価値であること", 6, false},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := newFranc(5).equals(newFranc(test.in)); got != test.want {
				t.Errorf("Franc.equals(): got %v want %v", got, test.want)
			}
		})
	}
}
