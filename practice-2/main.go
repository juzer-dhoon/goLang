package main

import (
	"fmt"

	"learning.com/practice2/custompackage"
)

var AppName = "GoCalc"
var version = "v1.0"

func main() {
	fmt.Println(AppName, version)

	var x, y float64
	var operator, name string

	fmt.Print("Enter your Name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter first number: ")
	fmt.Scanln(&x)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&y)

	result := custompackage.Calculate(x, y, operator)
	fmt.Println("Result:", result)
	fmt.Printf("Thank you for using %s, %s \n", AppName, name)
}
