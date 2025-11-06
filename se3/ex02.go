package main

import (
	"bufio"
	"fmt"
	"os"
)

// EXERCICE 12 : Lecture de fichier
// Créez une fonction qui lit un fichier et compte les lignes
// Input: chemin (string)
// Output: nombre de lignes (int), erreur

func compterLignes(chemin string) (int, error) {
	f, err := os.Open(chemin)
	ln := 0
	if err != nil {
		return 0, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ln++
	}
	return ln, nil
}

func main() {
	f, err := compterLignes("text.txt")

	if err != nil {
		fmt.Println("Ah oui, un problème occured:", err)
	}
	fmt.Println("Le nombre de ligne est de :",f)	
}