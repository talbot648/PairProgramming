package main

import (
	"errors"
	"math"
)

type TaxInfo struct {
	startOfTaxBand float64
	endOfTaxBand   float64
	taxRate        float64
}

var taxInformation = []TaxInfo{
	{startOfTaxBand: 1000000.00, endOfTaxBand: math.Inf(1), taxRate: .15},
	{startOfTaxBand: 750000.00, endOfTaxBand: 1000000.00, taxRate: .12},
	{startOfTaxBand: 325000.00, endOfTaxBand: 750000.00, taxRate: .10},
	{startOfTaxBand: 250000.00, endOfTaxBand: 325000.00, taxRate: .05},
	{startOfTaxBand: 145000.00, endOfTaxBand: 250000.00, taxRate: .02},
	{startOfTaxBand: 0, endOfTaxBand: 145000.00, taxRate: .00},
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
	totalTax := totalCurrentBandTax + getTaxFromPreviousBands(t.startOfTaxBand)
	return math.Floor(totalTax)
}

func getTaxFromPreviousBands(givenStartOfTaxBand float64) float64 {
	var totalTaxFromPreviousBands float64

	for _, band := range taxInformation {
		if band.endOfTaxBand <= givenStartOfTaxBand {
			taxableAmount := band.endOfTaxBand - band.startOfTaxBand
			bandTax := taxableAmount * band.taxRate
			totalTaxFromPreviousBands += bandTax
		}
	}
	return totalTaxFromPreviousBands
}

func getTaxBand(housePrice float64) TaxInfo {

	for _, taxBand := range taxInformation {
		if housePrice > taxBand.startOfTaxBand {
			return taxBand
		}
	}
	panic("couldn't find a tax band")

}

func isPriceValid(housePrice float64) bool {
	return housePrice > 0
}

func main() {
}
