package figure

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {

	areaTests := []struct {
		figure Figure
		want   float64
		err    error
	}{
		{Square{10}, 100, nil},
		{Square{-1}, 0, ErrInputSmall},
		{Circle{10}, 10 * 10 * math.Pi, nil},
		{Circle{0}, 0, ErrInputSmall},
	}

	assert := assert.New(t)
	for i, tt := range areaTests {
		got, err := tt.figure.Area()
		assert.Equal(got, tt.want, "[%d]got %.2f wanted %.2f", i, got, tt.want)
		assert.Equal(err, tt.err, "[%d]errors should be equal", i)
	}

}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		figure Figure
		want   float64
		err    error
	}{
		{Square{10}, 40, nil},
		{Square{-1}, 0, ErrInputSmall},
		{Circle{10}, 2 * math.Pi * 10, nil},
		{Circle{0}, 0, ErrInputSmall},
	}

	assert := assert.New(t)
	for i, tt := range perimeterTests {
		got, err := tt.figure.Perimeter()
		assert.Equal(got, tt.want, "[%d]got %.2f wanted %.2f", i, got, tt.want)
		assert.Equal(err, tt.err, "[%d]errors should be equal", i)
	}
}
