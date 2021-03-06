package main

import (
	"fmt"
	"math"
)

func (vec TwoDimVec) Coordinates() string {
	return fmt.Sprintf("(%v, %v)", vec.x, vec.y)
}

func (vec TwoDimVec) Magnitude() float64 {
	return math.Sqrt(math.Pow(vec.x, 2) + math.Pow(vec.y, 2))
}

func (vec TwoDimVec) TwoDimAngle() Angle {
	if vec.x == 0 {
		return Angle{math.Pi * 90 / 180, 90}
	}

	angle := math.Atan(vec.y / vec.x)

	return Angle{angle, angle * 180 / math.Pi}
}

func (vec TwoDimVec) Add(b TwoDimVec) TwoDimVec {
	return TwoDimVec{
		"C (A + B)",
		vec.x + b.x,
		vec.y + b.y,
	}
}

func (vec TwoDimVec) Sub(b TwoDimVec) TwoDimVec {
	return TwoDimVec{
		"D (A - B)",
		vec.x - b.x,
		vec.y - b.y,
	}
}

func (vec TwoDimVec) DotProduct(vecB TwoDimVec) float64 {
	product := vec.x*vecB.x + vec.y*vecB.y

	return product
}

func (vec TwoDimVec) CrossProduct(vecB TwoDimVec) float64 {
	product := vec.x*vecB.y - vec.y*vecB.x

	return product
}

func (vec TwoDimVec) AngleBetween(vecB TwoDimVec) Angle {
	magA := vec.Magnitude()
	magB := vecB.Magnitude()

	if magA == 0 && magB == 0 {
		return Angle{0, 0}
	}

	product := vec.DotProduct(vecB)
	angle := product / (magA * magB)

	return Angle{
		angle,
		angle * 180 / math.Pi,
	}
}
