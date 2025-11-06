package main

import "fmt"

// EXERCICE 4 : Tableaux et slices
// Créez une fonction qui trouve le maximum dans un slice
// Input: nombres []int
// Output: max (int)

func trouverMax(nombres []int) int {
	max := 0
	for _, elem := range nombres {
		if max < elem {
			max = elem
		}
	}
	return max
}

// // Version 4 : Moderne avec slices.Max (Go 1.21+)
// import "slices"

// func trouverMax(nombres []int) int {
// 	return slices.Max(nombres)  // Panic si vide
// }

func main() {
	n := []int{10,15,75,21,01,255,120}
	max := trouverMax(n)
	fmt.Println(max)
}

