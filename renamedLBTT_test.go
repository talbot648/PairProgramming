package main

import (
	"errors"
	"testing"
)

func TestReportsErrorWhenGivenInputBelowZero(t *testing.T) {
	givenPrice := -50000

	want := errors.New("invalid input: cannot have a house price at zero pounds or below")
	_, got := CalculateLBTT(givenPrice)

	if got.Error() != want.Error() {
		t.Error("expected error when given a house price below zero pounds")
	}
}

func TestReportsErrorWhenGivenInputisZero(t *testing.T) {
	givenPrice := 0

	want := errors.New("invalid input: cannot have a house price at zero pounds or below")
	_, got := CalculateLBTT(givenPrice)

	if got.Error() != want.Error() {
		t.Error("expected error when given a house price at zero pounds or below")
	}
}

func TestCalculatesZeroTaxForBelowFirsTaxBand(t *testing.T) {
	givenPrice := 100000

	want := 0
	got, gotErr := CalculateLBTT(givenPrice)

	if gotErr != nil {
		t.Error("unexpected error:", gotErr)
	}
	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}

}

func TestCalculatesZeroTaxAtMaximumValueBelowFirstTaxBand(t *testing.T) {
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

func TestCalculatesSecondTaxBand(t *testing.T) {
	givenPrice := 310000

	want := 5100
	got, _ := CalculateLBTT(givenPrice)

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}

}

func TestCalculateThirdTaxBand(t *testing.T) {
	givenPrice := 360000

	want := 9350

	got, _ := CalculateLBTT(givenPrice)

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func TestCalculateHighestTaxBand(t *testing.T) {
	givenPrice := 1400000

	want := 126350

	got, _ := CalculateLBTT(givenPrice)

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func TestAcceptanceTests(t *testing.T) {
	tests := []struct {
		name        string
		housePrice  int
		expected    int
		expectedErr error
	}{
		{name: "Invalid House Price Below Zero", housePrice: -5, expected: 0, expectedErr: errors.New("invalid input: cannot have a house price at zero pounds or below")},
	}

	for _, test := range tests {

		got, err := CalculateLBTT(test.housePrice)

		if test.expectedErr != nil {
			if err == nil {
				t.Errorf("%s (%v) Expected error, but received nil", test.name, test.housePrice)
			} else if err.Error() != test.expectedErr.Error() {
				t.Errorf("%s (%v) got error %v, want %v", test.name, test.housePrice, err, test.expectedErr)
			}
		} else if got != test.expected {
			t.Errorf("%s (%v) got %v, want %v", test.name, test.housePrice, got, test.expected)
		}
	}
}
