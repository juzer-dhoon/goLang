package custompackage

import "fmt"

func PrintMsg() {
	var text1 string = "Learning Go"
	var text2 int = 25
	var text3 float64 = 25.5
	fmt.Println("Someone is", text1, "Checking all variables: ", text2)
	fmt.Println("Checking all variables:")
	fmt.Println("Int: ", text2)
	fmt.Println("Float: ", text3)
	fmt.Println("String: ", text1)
	fmt.Println("Calling custom package.")
}
