package main 

func coefVector(v Vector, coe float64) Vector {
	return Vector{v.x * coe, v.y * coe, v.z * coe}
}
