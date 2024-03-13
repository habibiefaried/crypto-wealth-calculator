package main

import (
	"fmt"
	"math/big"
)

// ConvertStringToFloat shifts the decimal place of a large number in string format and returns it as a float64.
// The function takes two parameters: the number as a string and the number of decimal places to shift.
func ConvertStringToFloat(numStr string, decimalShift int) (float64, error) {
	// Convert the string to a big.Int
	numBigInt, success := new(big.Int).SetString(numStr, 10) // base 10
	if !success {
		return 0, fmt.Errorf("failed to convert string to big.Int")
	}

	// Create a big.Float from big.Int for division
	numBigFloat := new(big.Float).SetInt(numBigInt)

	// Calculate the divisor based on the decimalShift, which is effectively 10^decimalShift
	divisor := new(big.Float).SetFloat64(float64(1))
	for i := 0; i < decimalShift; i++ {
		divisor.Mul(divisor, big.NewFloat(10))
	}

	// Perform the division
	result := new(big.Float).Quo(numBigFloat, divisor)

	// Convert the result to a float64
	resultFloat, _ := result.Float64()

	return resultFloat, nil
}
