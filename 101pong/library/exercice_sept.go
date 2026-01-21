package main

import (
	"fmt"
	"math"
)

func hitAngle(v Vector) (float64, error) {
	norm_v := vectNorm(v)
	if norm_v == 0 {
		return 0.0, fmt.Errorf("The norm is nil")
	}
	angle_rad := math.Asin(math.Abs(v.z) / norm_v)
	return angle_rad * (180.0/math.Pi), nil
}