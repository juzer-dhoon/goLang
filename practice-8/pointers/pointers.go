package pointers

import "fmt"

// func valUpdate(x int) {
// 	x = 0
// }

func valUpdateAdress(x *int) {
	*x = 0
}

func ValUpdate() {
	i := 1
	fmt.Println("Simple Print:", i)

	// valUpdate(i)
	// fmt.Println("Value update: ", i)

	valUpdateAdress(&i)
	fmt.Println("Updated via Pointer: ", i)
}

func SumVariadic(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, val := range nums {
		total = total + val
	}
	fmt.Println("Total: ", total)
}
