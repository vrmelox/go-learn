package main

import "fmt"

// EXERCICE 5 : Maps
// Créez une fonction qui compte les occurrences de chaque mot
// Input: mots []string
// Output: map[string]int

func compterMots(mots []string) map[string]int {
	words := make(map[string]int)
	for _, elem := range mots {
		words[elem] += 1
	}
	return words
}

func main() {
	words := []string{"Génèse", "Exode", "Nombres", "Psaumes", "Génèse", "Psaumes", "Luc", "Luc", "Marc", "Psaumes", "Nombres", "Exode", "Apocalypse"}
	//roma := []string{}
	maping := compterMots(words)
	fmt.Println(maping)
}