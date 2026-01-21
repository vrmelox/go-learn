package main

import (
	"math"
)

func vectNorm(v Vector) float64 {
	return math.Sqrt((v.x * v.x) + (v.y * v.y) + (v.z * v.z))
}