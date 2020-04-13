package main

import "fmt"

type Point struct {
	x, y int
}

// This structure describes a square:
// start - coordinates of the upper left point of the square
// a - side length of a square
type Square struct {
	start Point
	a     uint
}

// Method which return perimeter of a square
func (p Square) Perimeter() uint {
	return p.a * 4
}

//  Method which return the coordinates of the lower right point of a square
func (p Square) End() Point {
	return Point{p.start.x + int(p.a), p.start.y - int(p.a)}
}

// Method which return area of a square
func (p Square) Area() uint {
	return p.a * p.a
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
