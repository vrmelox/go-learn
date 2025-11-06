package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// EXERCICE 14 : Manipulation de JSON
// Créez une fonction qui convertit une structure en JSON
// et une autre qui fait l'inverse

type Etudiant struct {
	Nom    string `json:"nom"`
	Age    int    `json:"age"`
	Notes  []int  `json:"notes"`
}

func etudiantVersJSON(e Etudiant) (string, error) {
	etue, err := json.MarshalIndent(e, ""," ")
	if err != nil {
		return "", errors.New("Conversion impossible")
	}
	return string(etue), nil
}

// func jsonVersEtudiant(donnees string) (Etudiant, error) {
// 	// VOTRE CODE ICI
// 	// Utilisez json.Unmarshal
// 	return Etudiant{}, nil
// }

func main() {
	etue := Etudiant{Nom: "Philippe", Age: 29, Notes: []int{17, 15, 20, 19, 18, 20}}
	etu, err := etudiantVersJSON(etue)
	if err != nil {
		fmt.Println("Houston, Houston, nous avons un problème:", err)
	}
	fmt.Println(etu)
}