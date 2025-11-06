package main 

import "fmt"

// EXERCICE 3 : Boucles
// Créez une fonction qui calcule la somme de 1 à n
// Input: n (int)
// Output: somme (int)

func sommeJusqua(n int) int {
	var result = 0
	for i := 1; i <= n ; i++ {
		result += i
	}
	return result
}

/*
func sommeJusqua(n int) int {
	result := 0
	for i := range n {  // Go 1.22+
		result += i + 1
	}
	return result
}
*/
func main() {
	var resultat = sommeJusqua(12)
	fmt.Println(resultat)
	fmt.Println(12 * 13)
}