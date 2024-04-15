package main

import "fmt"

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// Method to calculate area for Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Method to calculate perimeter for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Method to calculate area for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func areaCalculator() {
	rect := Rectangle{width: 10, height: 5}
	circle := Circle{radius: 5}

	// Call Area and Perimeter methods for Rectangle
	fmt.Println("Rectangle Area:", rect.Area())
	fmt.Println("Rectangle Perimeter:", rect.Perimeter())

	// Call Area method for Circle
	fmt.Println("Circle Area:", circle.Area())
}
