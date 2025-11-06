// EXERCICE 1 : Variables et types
// Créez une fonction qui calcule l'aire d'un rectangle
// Input: longueur, largeur (float64)
// Output: aire (float64)

package main

import "fmt"

func aireRectangle(longueur, largeur float64) float64 {
	return longueur * largeur
}

func main() {
	var aire float64 = aireRectangle(14.5, 7.56)
	ae := fmt.Sprintf("%.2f", aire)
	fmt.Println(ae)
}