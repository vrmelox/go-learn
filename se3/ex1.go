package main

import (
	"fmt"
	"math"
)

// EXERCICE 11 : Gestion d'erreurs personnalisées
// Créez un type d'erreur personnalisé pour "valeur négative"
// Créez une fonction CalculerRacineCarree qui retourne erreur si négatif

type ErreurValeurNegative struct {
	Valeur float64
}

func (e ErreurValeurNegative) Error() string {
	return fmt.Sprintf("%v est une valeur négative", e.Valeur)
}

func CalculerRacineCarree(x float64) (float64, error) {
	if x < 0 {
		return 0, ErreurValeurNegative{Valeur: x}
	}
	return math.Sqrt(x), nil
}

func main() {
	// Test avec valeur positive
	resultat, err := CalculerRacineCarree(16)
	if err != nil {
		fmt.Println("Erreur:", err)
	} else {
		fmt.Printf("√16 = %.2f\n", resultat)
	}
	
	// Test avec valeur négative
	resultat, err = CalculerRacineCarree(-25)
	if err != nil {
		fmt.Println("Erreur:", err)
		
		// Type assertion pour accéder aux détails
		if errNeg, ok := err.(ErreurValeurNegative); ok {
			fmt.Printf("Valeur problématique: %.2f\n", errNeg.Valeur)
		}
	} else {
		fmt.Printf("√-25 = %.2f\n", resultat)
	}
	
	// Test avec zéro
	resultat, err = CalculerRacineCarree(0)
	if err != nil {
		fmt.Println("Erreur:", err)
	} else {
		fmt.Printf("√0 = %.2f\n", resultat)
	}
}

/*
type ErreurCalcul struct {
	Valeur  float64
	Message string
}

func (e ErreurCalcul) Error() string {
	return fmt.Sprintf("erreur de calcul: %s (valeur: %.2f)", e.Message, e.Valeur)
}

func CalculerRacineCarree(x float64) (float64, error) {
	if x < 0 {
		return 0, ErreurCalcul{
			Valeur:  x,
			Message: "valeur négative",
		}
	}
	if math.IsNaN(x) {
		return 0, ErreurCalcul{
			Valeur:  x,
			Message: "valeur NaN (Not a Number)",
		}
	}
	if math.IsInf(x, 0) {
		return 0, ErreurCalcul{
			Valeur:  x,
			Message: "valeur infinie",
		}
	}
	return math.Sqrt(x), nil
}
*/

/*

*/