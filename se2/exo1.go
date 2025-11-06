package main

import (
	"fmt"
	"errors"
)

// EXERCICE 6 : Fonctions avec retours multiples
// Créez une fonction qui divise deux nombres
// Retourne le résultat et une erreur si division par zéro
// Input: a, b (float64)
// Output: résultat (float64), erreur (error)

func diviser(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, errors.New("division par zéro impossible")
	}
	return a/b, nil
}

func main() {
	res, err := diviser(12.7, 0)
	if err != nil {
		fmt.Println("Erreur:", err)
	} else {
		fmt.Printf("%v\n", res)
	}
}