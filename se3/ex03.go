package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// EXERCICE 13 : Écriture de fichier
// Créez une fonction qui écrit un slice de strings dans un fichier
// Input: chemin (string), lignes ([]string)
// Output: erreur

func ecrireLignes(chemin string, lignes []string) error {
	f, err := os.Create(chemin)
	if err != nil {
		return errors.New("Creation de fichier impossible")
	}
	defer f.Close()
	writer := bufio.NewWriter(f)

	for _, line := range lignes {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return errors.New("Problème survenue dans l'écriture des lignes")
		}
	}
	err = writer.Flush()
	if err != nil{
		errors.New("Toutes les données n'ont pas pu être écrites")
	}
	return nil
}

func main() {
	data := []string{"Go", "Down", "Low"}
	err := ecrireLignes("../danse.txt", data)
	if err != nil {
		fmt.Println("Un souci occured:", err)
		return
	}
}