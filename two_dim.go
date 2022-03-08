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

func (vec TwoDimVec) TwoDimAngle() TwoDimAngle {
	if vec.x == 0 {
		return TwoDimAngle{math.Pi * 90 / 180, 90}
	}

	angle := math.Atan(vec.y / vec.x)

	return TwoDimAngle{angle, angle * 180 / math.Pi}
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
