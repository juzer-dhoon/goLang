package main

import (
	"fmt"

	"learning.com/practice6/calculation"
)

func readInput() (float64, float64, float64) {
	var principal, rate, time float64

	fmt.Print("Enter Principal value: ")
	fmt.Scanln(&principal)

	fmt.Print("Enter Rate of interest: ")
	fmt.Scanln(&rate)

	fmt.Print("Enter Time (in years): ")
	fmt.Scanln(&time)

	return principal, rate, time
}

func main() {
	principal, rate, time := readInput()
	interest := calculation.CalculateInterest(principal, rate, time)
	fmt.Printf("Simple Interest: %.2f\n", interest)
}
