package main

import (
	"fmt"
	"maps"

	"learning.com/practice10/mapping"
)

func main() {
	// arr := make(map[string]int)

	// arr["key1"] = 50
	// arr["key2"] = 50

	// fmt.Println("array map initial values are :", arr)
	// fmt.Println("array map Calling specific values :", arr["key2"])
	// fmt.Println("Length of array :", len(arr))

	// checkVal := arr["key3"]
	// fmt.Println("Check value :", checkVal)
	// fmt.Println("array map values before assigning are :", arr)

	// arr["key3"] = 30
	// fmt.Println("array map values after assigning are :", arr)

	// delete(arr, "key3")
	// fmt.Println("array map values after deleting are :", arr)
	// fmt.Println("Length of array :", len(arr))

	// clear(arr)
	// fmt.Println("array map values after clearing are :", arr)
	// fmt.Println("Length of array :", len(arr))

	arrOne := map[string]int{"key1": 10, "key2": 20, "key3": 30}
	fmt.Println("Array :", arrOne)

	arrTwo := map[string]int{"key1": 10, "key2": 0, "key3": 30}
	fmt.Println("Array :", arrTwo)

	if maps.Equal(arrOne, arrTwo) {
		fmt.Println("Equals")
	} else {
		fmt.Println("Not Equals")
	}

	mapping.CalculateResult()
}
