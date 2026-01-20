package main

type Vector struct {
	x float64
	y float64
	z float64
}

func vectorize(x, y, z float64) Vector {
	return Vector{x, y, z}
}
