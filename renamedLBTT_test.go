package main

import (
	"errors"
	"testing"
)

func TestReportsErrorWhenGivenInputBelowZero(t *testing.T) {
	givenPrice := -50000.00

	want := errors.New("invalid input: cannot have a house price at zero pounds or below")
	_, got := CalculateLBTT(givenPrice)

	if got.Error() != want.Error() {
		t.Error("expected error when given a house price below zero pounds")
	}
}

func TestReportsErrorWhenGivenInputisZero(t *testing.T) {
	givenPrice := 0.00

	want := errors.New("invalid input: cannot have a house price at zero pounds or below")
	_, got := CalculateLBTT(givenPrice)

	if got.Error() != want.Error() {
		t.Error("expected error when given a house price at zero pounds or below")
	}
}

func TestCalculatesZeroTaxForBelowFirsTaxBand(t *testing.T) {
	givenPrice := 100000.00

	want := 0.00
	got, gotErr := CalculateLBTT(givenPrice)

	if gotErr != nil {
		t.Error("unexpected error:", gotErr)
	}
	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}

}

func TestCalculatesZeroTaxAtMaximumValueBelowFirstTaxBand(t *testing.T) {
	givenPrice := 145000.00

	want := 0.00
	got, gotErr := CalculateLBTT(givenPrice)

	if gotErr != nil {
		t.Error("unexpected error:", gotErr)
	}
	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}
func TestCalculatesFirstTaxBand(t *testing.T) {
	givenPrice := 200000.00

	want := 1100.00
	got, _ := CalculateLBTT(givenPrice)

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func TestCalculatesFirstTaxBandUpperEdgeCase(t *testing.T) {
	givenPrice := 250000.00

	want := 2100.00

	got, _ := CalculateLBTT(givenPrice)

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func TestCalculatesSecondTaxBand(t *testing.T) {
	givenPrice := 310000.00

	want := 5100.00
	got, _ := CalculateLBTT(givenPrice)

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}

}

func TestCalculateThirdTaxBand(t *testing.T) {
	givenPrice := 360000.00

	want := 9350.00

	got, _ := CalculateLBTT(givenPrice)

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func TestCalculateHighestTaxBand(t *testing.T) {
	givenPrice := 1400000.00

	want := 126350.00

	got, _ := CalculateLBTT(givenPrice)

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func TestRoundNumberToTwoDecimalPlaces(t *testing.T) {
	givenPrice := 768549.64
	want := 50575.96

	got, _ := CalculateLBTT(givenPrice)
	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func TestAcceptanceTests(t *testing.T) {
	tests := []struct {
		name        string
		housePrice  float64
		expected    float64
		expectedErr error
	}{
		{name: "Invalid House Price Below Zero", housePrice: -5.00, expected: 0, expectedErr: errors.New("invalid input: cannot have a house price at zero pounds or below")},
		{name: "Invalid House Price at Zero", housePrice: 0.00, expected: 0, expectedErr: errors.New("invalid input: cannot have a house price at zero pounds or below")},
		{name: "Zero Tax below first tax band", housePrice: 110000.00, expected: 0, expectedErr: nil},
		{name: "First Tax Band", housePrice: 150000.00, expected: 100.00, expectedErr: nil},
		{name: "Second Tax Band", housePrice: 300000.00, expected: 4600.00, expectedErr: nil},
		{name: "Third Tax Band", housePrice: 400000.00, expected: 13350.00, expectedErr: nil},
		{name: "Highest Tax Band", housePrice: 800000.00, expected: 54350.00, expectedErr: nil},
		{name: "Edge Case at First Tax Band Start", housePrice: 145000.00, expected: 0, expectedErr: nil},
		{name: "Edge Case at Second Tax Band Start", housePrice: 250000.00, expected: 2100.00, expectedErr: nil},
		{name: "Edge Case at Third Tax Band Start", housePrice: 325000.00, expected: 5850.00, expectedErr: nil},
		{name: "Edge Case at Highest Tax Band Start", housePrice: 750000.00, expected: 48350.00, expectedErr: nil},
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
