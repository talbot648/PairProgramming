package main

import (
	"errors"
)

func main() {
}

func CalculateLBTT(housePrice int) (int, error) {
	if !isPriceValid(housePrice) {
		return 0, errors.New("invalid input: cannot have a house price below zero pounds")
	}
	return 0, errors.New("")
}

func isPriceValid(housePrice int) bool {
	return housePrice >= 0
}
