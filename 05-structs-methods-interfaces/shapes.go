package shapes

import "math"

type Rectangle struct {
	x float64
	y float64
}

type Circle struct {
	Radius float64
}

type Traingle struct {
	Base   float64
	Height float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.x + rect.y)
}

func (r Rectangle) Area() float64 {
	return r.x * r.y
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Traingle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
