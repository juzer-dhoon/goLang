package main

import "fmt"

func main() {
	// initializing array
	a := [5]int{} //short notation
	fmt.Println("Value of a is :", a)

	// assigning value to array
	b := [5]int{5, 6, 7, 8, 9}
	b[2] = 100
	fmt.Println("Value of b is :", b)

	// two dimensional array
	xyz := [5][3]int{}

	// assigning value two dimensional array
	xyz[0][1] = 10
	xyz[0][2] = 20
	xyz[1][2] = 30
	fmt.Println("Value of xyz is :", xyz)
	fmt.Println("Length of xyz is :", len(xyz))

	// string array
	at := [4]string{"Dimple", "Ishan", "Kalpesh", "Dharmesh"}
	at[0] = "Juzer"
	fmt.Println("Value of AT is :", at)

}
