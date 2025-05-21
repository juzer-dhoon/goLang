package main

import (
	"fmt"

	"learning.com/practice9/recursion"
)

func CloseureIntSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	newSeq := CloseureIntSeq()
	fmt.Println(newSeq())
	fmt.Println(newSeq())
	fmt.Println(newSeq())
	fmt.Println(newSeq())
	fmt.Println(newSeq())
	fmt.Println(newSeq())

	var x int
	fmt.Print("Please enter a number to get a factorial: ")
	fmt.Scan(&x)

	result := recursion.Fact(x)
	fmt.Printf("Factorial of %d is %d \n", x, result)

	z := 28
	result2 := recursion.SumOfDigit(z)
	fmt.Println("The total sum is : ", result2)

	str := "Juzer"
	revstr := recursion.RevString(str)
	fmt.Println("The rev string is : ", revstr)

	// str := "hello"
	// fmt.Println(str[0])
	// fmt.Println(string(str[0]))
}
