package main

import (
	"errors"
)

/*
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

*/

type TaxInfo struct {
	startOfTaxBand            int
	taxRate                   int
	totalTaxFromPreviousBands int
}

var taxInformation = []TaxInfo{
	{startOfTaxBand: 0, taxRate: 0, totalTaxFromPreviousBands: 0},
	{startOfTaxBand: 145000, taxRate: 2, totalTaxFromPreviousBands: 0},
	{startOfTaxBand: 250000, taxRate: 5, totalTaxFromPreviousBands: 2100},
	{startOfTaxBand: 325000, taxRate: 10, totalTaxFromPreviousBands: 5850},
	{startOfTaxBand: 750000, taxRate: 12, totalTaxFromPreviousBands: 48350},
}

func CalculateLBTT(housePrice int) (int, error) {
	var totalTax int
	if !isPriceValid(housePrice) {
		return 0, errors.New("invalid input: cannot have a house price at zero pounds or below")
	}
	if housePrice <= 145000 {
		return totalTax, nil
	}
	taxBand := getTaxBandIndex(housePrice)
	return taxBand.calculateTax(housePrice), nil
	/*
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
	*/

}

func main() {
}

func (t *TaxInfo) calculateTax(housePrice int) int {
	totalToBeTaxed := housePrice - t.startOfTaxBand

	totalCurrentBandTax := totalToBeTaxed * t.taxRate / 100
	totalTax := totalCurrentBandTax + t.totalTaxFromPreviousBands

	return totalTax
}

func getTaxBandIndex(housePrice int) TaxInfo {
	for i := len(taxInformation) - 1; i >= 0; i-- {
		if housePrice > taxInformation[i].startOfTaxBand {
			return taxInformation[i]
		}
	}
	return taxInformation[0] // Default to the first band if no match is found
}

func isPriceValid(housePrice int) bool {
	return housePrice > 0
}

/*

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
*/
