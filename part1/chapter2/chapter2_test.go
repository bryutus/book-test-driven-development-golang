package chapter2

import (
	"testing"
)

func TestMultiplication(t *testing.T) {

	testCases := []struct {
		desc string
		in   int
		want int
	}{
		{"$5*2=10がとなること", 2, 10},
		{"$5*3=15がとなること", 3, 15},
	}

	five := newDollar(5)

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := five.times(test.in); got.Amount != test.want {
				t.Errorf("Doller.times(): got.Amout %v want %v", got.Amount, test.want)
			}
		})
	}
}
