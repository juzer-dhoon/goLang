package main

import (
	"fmt"

	"learning.com/practice8/pointers"
)

func printEvenNumber(x int) {
	for n := 0; n < x; n++ {
		if n%2 == 0 {
			fmt.Println("Even numbers: ", n)
		}
	}
	// return x
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for j := range 5 {
		fmt.Println(j)
	}

	for n := range 10 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("Check: ", n)
	}

	for {
		fmt.Println("Checking Loop")
		break
	}

	var limit int
	fmt.Print("Enter the number: ")
	fmt.Scanln(&limit)
	printEvenNumber(limit)

	pointers.ValUpdate()
	pointers.SumVariadic(2, 5, 6, 8)
}
