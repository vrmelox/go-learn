package main

import "fmt"

// EXERCICE 2 : Conditions
// Créez une fonction qui retourne "pair" ou "impair"
// Input: nombre (int)
// Output: string

func pairOuImpair(n int) string {
	if n % 2 == 0 {
		return "Pair"
	}
	return "Impair"
}

func main() {
	verdict := pairOuImpair(40)
	fmt.Println(verdict)
}
