package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	var input float64

	fmt.Print("Enter a number to divide: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Invalid input for number")
		os.Exit(1)
	}

	original := big.NewFloat(input)
	dividedValues := []*big.Float{}
	sum := big.NewFloat(0)
	ratio := big.NewFloat(2)
	// Precision threshold to stop dividing when values get too small
	threshold := big.NewFloat(1e-6)
	previousValue := new(big.Float).Set(original)    

	fmt.Println("\n--- Dividing Values ---")
	for i := 1; ; i++ {
		divided := new(big.Float).Quo(previousValue, ratio)
		if divided.Cmp(threshold) <= 0 {
			break
		}
		dividedValues = append(dividedValues, divided)
		sum.Add(sum, divided)
		fmt.Printf("Step %2d: %s\n", i, divided.Text('f', 10))   
		if sum.Cmp(original) >= 0 {
			break
		}

		previousValue = divided
	}

	lastIndex := len(dividedValues) - 1
	dividedValues[lastIndex].Sub(dividedValues[lastIndex], new(big.Float).Sub(sum, original))
	fmt.Println("\n--- Results ---")
	finalSum := big.NewFloat(0)
	for _, val := range dividedValues {
		finalSum.Add(finalSum, val)
	}

	fmt.Printf("Final Sum:    %s\n", finalSum.Text('f', 10))
	fmt.Printf("Original Value: %s\n\n", original.Text('f', 10))
}
