package main

import (
	"errors"
	"math"
)

type TaxInfo struct {
	startOfTaxBand            float64
	taxRate                   float64
	totalTaxFromPreviousBands float64
}

var taxInformation = []TaxInfo{
	{startOfTaxBand: 145000.00, taxRate: .02, totalTaxFromPreviousBands: 0.00},
	{startOfTaxBand: 250000.00, taxRate: .05, totalTaxFromPreviousBands: 2100.00},
	{startOfTaxBand: 325000.00, taxRate: .10, totalTaxFromPreviousBands: 5850.00},
	{startOfTaxBand: 750000.00, taxRate: .12, totalTaxFromPreviousBands: 48350.00},
}

func CalculateLBTT(housePrice float64) (float64, error) {
	if !isPriceValid(housePrice) {
		return 0, errors.New("invalid input: cannot have a house price at zero pounds or below")
	}

	if housePrice <= 145000 {
		return 0, nil
	}
	taxBand := getTaxBand(housePrice)
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

func (t *TaxInfo) calculateTax(housePrice float64) float64 {
	totalToBeTaxed := housePrice - t.startOfTaxBand

	totalCurrentBandTax := totalToBeTaxed * t.taxRate
	totalTax := totalCurrentBandTax + t.totalTaxFromPreviousBands
	return math.Round(totalTax*100) / 100
}

func getTaxBand(housePrice float64) TaxInfo {
	for i := len(taxInformation) - 1; i >= 0; i-- {
		if housePrice > taxInformation[i].startOfTaxBand {
			return taxInformation[i]
		}
	}
	return taxInformation[0]
}

func isPriceValid(housePrice float64) bool {
	return housePrice > 0
}

func main() {
}

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
