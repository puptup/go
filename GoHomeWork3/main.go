package main

import "fmt"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (p Square) Perimeter() uint {
	return p.a * 4
}

func (p Square) End() Point {
	return Point{p.start.x + int(p.a), p.start.y - int(p.a)}
}

func (p Square) Area() uint {
	return p.a * p.a
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
