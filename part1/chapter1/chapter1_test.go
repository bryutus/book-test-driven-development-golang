package chapter1

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := newDollar(5)
	five.times(2)

	if got, want := five.Amount, 10; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
