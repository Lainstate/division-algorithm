package main

import (
	"fmt"
	"math"
)

func divideUntilSumEqualsOriginal(n float64) {
	// List to store all divisions
	var divisions []float64
	current := n
	var sum float64
	epsilon := 1e-12 // Small threshold for floating-point precision

	for {
		// If the current value divided by 2 is less than epsilon, stop
		if current/2 < epsilon {
			divisions = append(divisions, current)
			sum += current
			break
		}

		// Divide by two
		half := current / 2
		divisions = append(divisions, half)
		sum += half

		// Update current for the next iteration
		current -= half

		// Stop if the difference between sum and n is within epsilon
		if math.Abs(n-sum) < epsilon {
			break
		}
	}

	// Print each division with high precision
	fmt.Println("Divisions:")
	for i, div := range divisions {
		fmt.Printf("Step %d: %.12f\n", i+1, div)
	}

	// Print the sum of divisions and compare it to the original number
	fmt.Printf("\nSum of divisions: %.12f\n", sum)
	fmt.Printf("Original number: %.12f\n", n)

	// Check if the sum is close enough to the original number
	if math.Abs(sum-n) < epsilon {
		fmt.Println("The sum of divisions equals the original number.")
	} else {
		fmt.Println("Precision error: The sum does not exactly match the original number.")
	}
}

func main() {
	var number float64
	fmt.Print("Enter a number (can be decimal): ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	divideUntilSumEqualsOriginal(number)
}
