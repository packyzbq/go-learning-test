package structs

import "math"

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Cycle struct
type Cycle struct {
	Radius float64
}

// Shape interface
type Shape interface {
	Area() float64
}

// Perimeter method
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area method
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter method
func (c Cycle) Perimeter() float64 {
	return math.Pi * c.Radius
}

// Area method
func (c Cycle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2.0)
}
