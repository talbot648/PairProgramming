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

func TestCalculatesZeroTaxForBelowFirsTaxBand(t *testing.T) {
	givenPrice := 145000

	want := 0
	got, gotErr := CalculateLBTT(givenPrice)

	if gotErr != nil {
		t.Error("unexpected error:", gotErr)
	}
	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}

}

func TestCalculatesFirstTaxBand(t *testing.T) {
	givenPrice := 200000

	want := 1100
	got, _ := CalculateLBTT(givenPrice)

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func TestCalculatesFirstTaxBandUpperEdgeCase(t *testing.T) {
	givenPrice := 250000

	want := 2100

	got, _ := CalculateLBTT(givenPrice)

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}
