package array

import (
	"fmt"
	"slices"
)

func matchArray() {
	firstArray := []int{1, 2, 3}
	SecondArray := []int{1, 2, 3}
	fmt.Println("Both array are::", slices.Equal(firstArray, SecondArray))
}

func PrintArray() {
	var sliceArray []string
	fmt.Println("Slice Array values: ", sliceArray)
	sliceArray = make([]string, 2, 5)
	fmt.Println("Slice Array values: ", sliceArray)
	fmt.Println("Length of Slice Array values: ", len(sliceArray))
	fmt.Println("Capacity of Slice Array values: ", cap(sliceArray))

	sliceArray = append(sliceArray, "P")
	sliceArray = append(sliceArray, "Q")
	sliceArray = append(sliceArray, "R")
	sliceArray = append(sliceArray, "S")

	fmt.Println("Slice Array values: ", sliceArray)
	fmt.Println("Length of Slice Array values: ", len(sliceArray))
	fmt.Println("Capacity of Slice Array values: ", cap(sliceArray))

	sliceArray[0] = "N"
	sliceArray[1] = "O"

	fmt.Println("Slice Array values: ", sliceArray)
	fmt.Println("Length of Slice Array values: ", len(sliceArray))
	fmt.Println("Capacity of Slice Array values: ", cap(sliceArray))

	copySliceArray := make([]string, len(sliceArray))
	copy(copySliceArray, sliceArray)
	fmt.Println("Copy of slice array: ", copySliceArray)
	copySliceArray[5] = "Z"
	fmt.Println("Copy of slice array and new added: ", copySliceArray)
	fmt.Println("Copy of slice array and printing limited after values: ", copySliceArray[:5])
	fmt.Println("Copy of slice array and printing limited before values: ", copySliceArray[2:])
	fmt.Println("Copy of slice array and new added: ", copySliceArray[2:5])
	// fmt.Println("Copy of slice array with reducing length: ", copySliceArray[len(copySliceArray)-2])
	fmt.Println("Copy of slice array except last two: ", copySliceArray[:len(copySliceArray)-2])
	matchArray()
}
