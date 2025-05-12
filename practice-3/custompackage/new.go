package custompackage

import "fmt"

func Add(x, y float64) float64 {
	return x + y
}

func Subtract(x, y float64) float64 {
	return x - y
}

func Multiply(x, y float64) float64 {
	return x * y
}

func Divide(x, y float64) float64 {
	if y != 0 {
		return x / y
	}
	fmt.Println("Error: Division by zero")
	return 0
}

func Calculate(x, y float64, operator string) float64 {
	if operator == "+" {
		return Add(x, y)
	} else if operator == "-" {
		return Subtract(x, y)
	} else if operator == "*" {
		return Multiply(x, y)
	} else if operator == "/" {
		return Divide(x, y)
	} else {
		fmt.Println("Invalid operator")
		return 0
	}
}
