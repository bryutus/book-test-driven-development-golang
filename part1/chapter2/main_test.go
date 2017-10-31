package main

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := newDollar(5)

	product := five.times(2)
	if got, want := product.Amount, 10; got != want {
		t.Errorf("got %v want %v", got, want)
	}

	product = five.times(3)
	if got, want := product.Amount, 15; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
