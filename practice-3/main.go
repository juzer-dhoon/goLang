package main

import (
	"bufio"
	"fmt"
	"os"

	"learning.com/practice3/custompackage"
)

var AppName = "GoCalc"
var version = "v1.0"

func readInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	fmt.Println(AppName, version)

	var x, y float64
	var operator, name string

	name = readInput("Enter your name: ")
	fmt.Println("Hello,", name)

	fmt.Print("Enter first number: ")
	fmt.Scanln(&x)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&y)

	result := custompackage.Calculate(x, y, operator)
	fmt.Println("Result:", result)
	fmt.Printf("Thank you for using %s, %s\n", AppName, name)
}
