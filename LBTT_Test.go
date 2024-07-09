package main

import (
	"errors"
	"testing"
)

func TestReportsErrorWhenGivenInputBelowZero(t *testing.T) {
	givenPrice := -1

	want := errors.New("invalid input: cannot have a house price below zero pounds")
	_, got := CalculateLBTT(givenPrice)

	if got.Error() != want.Error() {
		t.Error("expected error when given a house price below zero pounds")
	}
}
