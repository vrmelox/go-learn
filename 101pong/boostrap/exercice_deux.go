package main

func vectorSum(a, b Vector) Vector {
	return Vector{a.x + b.x, a.y + b.y, a.z + b.z}
}

func vectorDif(a, b Vector) Vector {
	return Vector{b.x - a.x, b.y - a.y, b.z - a.z}
}

// func main() {
// 	veca := vectorSum(Vector{4.0, 5.2, 1.7}, Vector{4.5, 0.8, 1.7})
// 	vecb := vectorDif(Vector{4.0, 5.2, 1.7}, Vector{4.5, 0.8, 1.7})
// 	fmt.Println("Le résultat de veca est de ", veca)
// 	fmt.Println("Le résultat de veca est de ", vecb)
// }	