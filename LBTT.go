package main

import (
	"errors"
)

func calculateFirstTaxBandTotal(housePrice int) int {
	startOfTaxBand := 145000
	taxRate := 2 //percent

	totalToBeTaxed := housePrice - startOfTaxBand
	totalTax := totalToBeTaxed * taxRate / 100

	return totalTax
}

func calculateSecondTaxBandTotal(housePrice int) int {
	startOfTaxBand := 250000
	taxRate := 5 //percent
	taxFromPreviousTaxBand := 2100

	totalToBeTaxed := housePrice - startOfTaxBand
	totalCurrentBandTax := totalToBeTaxed * taxRate / 100
	totalTax := totalCurrentBandTax + taxFromPreviousTaxBand

	return totalTax
}

func calculateThirdTaxBandTotal(housePrice int) int {
	startOfTaxBand := 325000
	taxRate := 10 //percent
	totalTaxFromPreviousBands := 5850

	totalToBeTaxed := housePrice - startOfTaxBand
	totalCurrentBandTax := totalToBeTaxed * taxRate / 100
	totalTax := totalCurrentBandTax + totalTaxFromPreviousBands

	return totalTax

}

func calculateHighestTaxBandTotal(housePrice int) int {
	startOfTaxBand := 750000
	taxRate := 12 //percent
	totalTaxFromPreviousBands := 48350

	totalToBeTaxed := housePrice - startOfTaxBand
	totalCurrentBandTax := totalToBeTaxed * taxRate / 100
	totalTax := totalCurrentBandTax + totalTaxFromPreviousBands

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
	if isPriceInSecondTaxBand(housePrice) {
		totalTax := calculateSecondTaxBandTotal(housePrice)
		return totalTax, nil
	}
	if isPriceInThirdTaxBand(housePrice) {
		totalTax := calculateThirdTaxBandTotal(housePrice)
		return totalTax, nil
	}
	if isPriceInHighestTaxBand(housePrice) {
		totalTax := calculateHighestTaxBandTotal(housePrice)
		return totalTax, nil
	}

	return 0, errors.New("")
}

func isPriceValid(housePrice int) bool {
	return housePrice >= 0
}

func isPriceInFirstTaxBand(housePrice int) bool {
	return housePrice > 145000 && housePrice <= 250000
}

func isPriceInSecondTaxBand(housePrice int) bool {
	return housePrice > 250000 && housePrice <= 325000
}

func isPriceInThirdTaxBand(housePrice int) bool {
	return housePrice > 325000 && housePrice <= 750000
}

func isPriceInHighestTaxBand(housePrice int) bool {
	return housePrice > 750000
}
