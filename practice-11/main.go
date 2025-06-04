package main

import (
	"fmt"

	"learning.com/practice11/structure"
)

type person struct {
	name string
	city string
	age  int
}

// func newPerson(name string, age int, city ...string) person {
// 	defaultCity := "Mumbai"
// 	if len(city) > 0 {
// 		return person{name: name, city: city[1], age: age}
// 	}
// 	return person{name: name, city: defaultCity, age: age}
// }

// func newPerson(name string, age int, city string) person {
// 	p := person{
// 		name: name,
// 		age:  age,
// 		city: city,
// 	}
// 	return p
// }

func main() {
	fmt.Println(person{"Juzer", "Surat", 30})
	fmt.Println(person{name: "Juzer"})
	fmt.Println(&person{name: "Juzer"})

	// p1 := newPerson("Juzer", 30, "Surat", "Vadodara")
	// fmt.Println(p1)

	// x := newPerson("JD", 31, "Surat")
	// fmt.Println(x)

	vehicle := struct {
		model     string
		isAvgGood bool
	}{
		"Chev",
		true,
	}
	fmt.Println(vehicle)

	structure.StudentResult()
}
