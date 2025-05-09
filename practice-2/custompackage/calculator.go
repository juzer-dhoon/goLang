package custompackage

import "fmt"

func Calculate(x, y float64, operator string) float64 {
	switch operator {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		if y != 0 {
			return x / y
		}
		fmt.Println("Error: Division by zero")
		return 0
	default:
		fmt.Println("Invalid operator")
		return 0
	}
}
