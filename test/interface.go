package figure

import (
	"errors"
	"math"
)

//Figure describes methods:
//area - area of figure
//perimeter - perimeter of figure
type Figure interface {
	Area() (float64, error)
	Perimeter() (float64, error)
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

var ErrInputSmall = errors.New("bad input. too small")

func (s Square) Area() (float64, error) {
	if s.side <= 0 {
		return 0, ErrInputSmall
	}
	return s.side * s.side, nil
}

func (s Square) Perimeter() (float64, error) {
	if s.side <= 0 {
		return 0, ErrInputSmall
	}
	return s.side * 4, nil
}

func (c Circle) Area() (float64, error) {
	if c.radius <= 0 {
		return 0, ErrInputSmall
	}
	return math.Pi * c.radius * c.radius, nil
}

func (c Circle) Perimeter() (float64, error) {
	if c.radius <= 0 {
		return 0, ErrInputSmall
	}
	return 2 * math.Pi * c.radius, nil
}
