package main

import "fmt"

type rect struct {
	width  float64
	height float64
}

type cir struct {
	radius float64
}

type Shape interface {
	area() float64
	perimeter() float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c cir) area() float64 {
	return 3.14 * c.radius * c.radius
}

func printDetails(s Shape) {
	fmt.Println("Area of rect: ", s.area())
	fmt.Println("Area of rect: ", s.perimeter())
}

func main() {
	r := rect{width: 20, height: 10}
	printDetails(r)

	// rectArea := r.area()
	// fmt.Println("Area of rect: ", rectArea)

	// rectPerimeter := r.perimeter()
	// fmt.Println("Perimeter of rect: ", rectPerimeter)

	c := cir{radius: 10}
	cirArea := c.area()
	fmt.Println("Area of circle: ", cirArea)

}
