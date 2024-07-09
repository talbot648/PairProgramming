package main

import (
	"errors"
)

func calculateFirstTaxBandTotal(housePrice int) int {
	startOfTaxBand := 145000
	taxRate := 2 //percent

	totalToBeTaxed := housePrice - startOfTaxBand
	totalTax := totalToBeTaxed / 100 * taxRate

	return totalTax
}

func main() {
}

func CalculateLBTT(housePrice int) (int, error) {
	var totalTax int
	if !isPriceValid(housePrice) {
		return 0, errors.New("invalid input: cannot have a house price below zero pounds")
	}
	if housePrice <= 145000 {
		return totalTax, nil
	}

	if isPriceInFirstTaxBand(housePrice) {
		totalTax := calculateFirstTaxBandTotal(housePrice)
		return totalTax, nil
	}

	return 0, errors.New("")
}

func isPriceValid(housePrice int) bool {
	return housePrice >= 0
}

func isPriceInFirstTaxBand(housePrice int) bool {
	return housePrice >= 145001 && housePrice <= 250000
}
