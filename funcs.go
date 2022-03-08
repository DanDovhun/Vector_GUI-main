package main

import "math"

func round(num float64, to int) float64 {
	return math.Round(num*math.Pow10(to)) / math.Pow10(to)
}
