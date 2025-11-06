package main

// EXERCICE 11 : Gestion d'erreurs personnalisées
// Créez un type d'erreur personnalisé pour "valeur négative"
// Créez une fonction CalculerRacineCarree qui retourne erreur si négatif

type ErreurValeurNegative struct {
	Valeur float64
}

func (e ErreurValeurNegative) Error() string {
	// VOTRE CODE ICI
	return ""
}

func CalculerRacineCarree(x float64) (float64, error) {
	// VOTRE CODE ICI
	return 0, nil
}