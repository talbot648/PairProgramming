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
	{startOfTaxBand: 1000000.00, taxRate: .15, totalTaxFromPreviousBands: 78350.00},
	{startOfTaxBand: 750000.00, taxRate: .12, totalTaxFromPreviousBands: 48350.00},
	{startOfTaxBand: 325000.00, taxRate: .10, totalTaxFromPreviousBands: 5850.00},
	{startOfTaxBand: 250000.00, taxRate: .05, totalTaxFromPreviousBands: 2100.00},
	{startOfTaxBand: 145000.00, taxRate: .02, totalTaxFromPreviousBands: 0.00},
	{startOfTaxBand: 0, taxRate: .00, totalTaxFromPreviousBands: 0.00},
}

func CalculateLBTT(housePrice float64) (float64, error) {
	if !isPriceValid(housePrice) {
		return 0, errors.New("invalid input: cannot have a house price at zero pounds or below")
	}

	taxBand := getTaxBand(housePrice)
	return taxBand.calculateTax(housePrice), nil

}

func (t *TaxInfo) calculateTax(housePrice float64) float64 {
	totalToBeTaxed := housePrice - t.startOfTaxBand

	totalCurrentBandTax := totalToBeTaxed * t.taxRate
	totalTax := totalCurrentBandTax + t.totalTaxFromPreviousBands
	return math.Floor(totalTax)
}

func getTaxBand(housePrice float64) TaxInfo {

	for _, taxBand := range taxInformation {
		if housePrice > taxBand.startOfTaxBand {
			return taxBand
		}
	}
	panic("couldn't find a tax band")
	/*
		for i := len(taxInformation) - 1; i >= 0; i-- {
			if housePrice > taxInformation[i].startOfTaxBand {
				return taxInformation[i]
			}
		}
		return taxInformation[0]
	*/
}

func isPriceValid(housePrice float64) bool {
	return housePrice > 0
}

func main() {
}
