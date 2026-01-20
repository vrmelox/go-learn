package main 

import (
	"math"
)

func vectNorm(v Vector) float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2) + math.Pow(v.z, 2))
}