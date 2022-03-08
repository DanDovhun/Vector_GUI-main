package main

import (
	"fmt"
	"math"
)

func (vec ThreeDimVec) Coordinates() string {
	return fmt.Sprintf("(%v, %v, %v)", vec.x, vec.y, vec.z)
}

func (vec ThreeDimVec) Magnitude() float64 {
	return math.Sqrt(math.Pow(vec.x, 2) + math.Pow(vec.y, 2) + math.Pow(vec.z, 2))
}

func (vec ThreeDimVec) Elevation() Angle {
	if vec.x == 0 && vec.z == 0 {
		return Angle{90 * math.Pi / 180, 90}
	}

	small := TwoDimVec{"A", vec.x, vec.z}
	mag := small.Magnitude()
	angle := math.Atan(vec.y / mag)

	return Angle{
		angle,
		180 * angle / math.Pi,
	}
}

func (vec ThreeDimVec) Rotation() Angle {
	if vec.x == 0 {
		return Angle{90 * math.Pi / 180, 90}
	}

	angle := math.Atan(vec.z / vec.x)

	return Angle{
		angle,
		180 * angle / math.Pi,
	}
}

func (vec ThreeDimVec) Add(b ThreeDimVec) ThreeDimVec {
	return ThreeDimVec{
		"C (A + B)",
		vec.x + b.x,
		vec.y + b.y,
		vec.z + b.z,
	}
}

func (vec ThreeDimVec) Sub(b ThreeDimVec) ThreeDimVec {
	return ThreeDimVec{
		"D (A - B)",
		vec.x + b.x,
		vec.y + b.y,
		vec.z + b.z,
	}
}

func (vec ThreeDimVec) DotProduct(vecB ThreeDimVec) float64 {
	product := vec.x*vecB.x + vec.y*vecB.y + vec.z*vecB.z

	return product
}

func (vec ThreeDimVec) CrossProduct(vecB ThreeDimVec) ThreeDimVec {
	x := vec.y*vecB.z - vec.z*vecB.y
	y := vec.x*vecB.z - vec.z*vecB.x
	z := vec.x*vecB.y - vec.y*vecB.x

	return ThreeDimVec{"E (Cross product)", x, y, z}
}

func (vec ThreeDimVec) AngleBetween(vecB ThreeDimVec) Angle {
	magA := vec.Magnitude()
	magB := vecB.Magnitude()

	if magA == 0 && magB == 0 {
		return Angle{0, 0}
	}

	dotProcuct := vec.DotProduct(vecB)
	angle := dotProcuct / (magA * magB)

	return Angle{
		angle,
		180 * angle / math.Pi,
	}
}
