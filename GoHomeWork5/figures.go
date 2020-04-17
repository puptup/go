package main

import (
	"fmt"
	"math"
)

//Figure describes methods:
//area - area of figure
//perimeter - perimeter of figure
type Figure interface {
	area() float64
	perimeter() float64
}

//Square structure describes square
type Square struct {
	//side length of a square
	side float64
}

//Circle structure describes circle
type Circle struct {
	//radius of circle
	radius float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (s Square) perimeter() float64 {
	return s.side * 4
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {

	// square with lengh of side equals 10 should return area = 100 , perimeter = 40
	var s Figure = Square{10}
	if s.area() != float64(100) {
		fmt.Print("area of square is incorrect")
	} else if s.perimeter() != float64(40) {
		fmt.Print("perimeter of square is incorrect")
	} else {
		fmt.Println("The square is counted correctly")
	}

	// circle with radius equals 5 should return area = pi * 25, perimeter = 10 * pi
	var c Figure = Circle{5}
	if c.area() != math.Pi*25 {
		fmt.Print("area of circle is incorrect")
	} else if c.perimeter() != 10*math.Pi {
		fmt.Print("perimeter of circle is incorrect")
	} else {
		fmt.Println("The circle is counted correctly")
	}

	fmt.Println(s.area(), s.perimeter())
	fmt.Println(c.area(), c.perimeter())

}
