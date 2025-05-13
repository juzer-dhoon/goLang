package custompage

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var PublicVar int = 100
var privateVar int = 200

func readInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func PublicMethod() {
	fmt.Println("Hello, this is Public Method")
	fmt.Println("Calling Public var:: ", PublicVar)
	privateMethod()
}

func privateMethod() {
	fmt.Println("Hello, this is Private Method")
	fmt.Println("Calling Private var:: ", privateVar)
}

func StringToFloat() {
	a := readInput("Enter enter string: ")
	b, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println("Error is coming :", err)
		return
	}
	fmt.Println(b)
	fmt.Printf("Type of b is %T \n", b)
}

func FloatToInt() {
	var a float64
	fmt.Print("Enter enter float value: ")
	fmt.Scanln(&a)
	b := int(a)
	fmt.Println(b)
}
