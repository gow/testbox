package interfaces

import "math"

type geometry interface {
	shape() string
	area() float64
	perimeter() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) shape() string {
	return "Square"
}
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) perimeter() float64 {
	return 4 * s.length
}

func (c circle) shape() string {
	return "Circle"
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) (shape string, area float64, peri float64) {
	shape = g.shape()
	area = g.area()
	peri = g.perimeter()
	return // Automatically returns (shape, area, peri)
}
