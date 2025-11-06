package main

import "fmt"

// EXERCICE 7 : Structures
// Créez une structure Personne avec nom, age
// Créez une méthode EstMajeur() qui retourne true si age >= 18

type Personne struct {
	Nom		string
	Age		int
}

func (p Personne) EstMajeur() bool {
	if p.Age >= 18 {
		return true
	}
	return false
}

func main() {
	elise := Personne{
		Nom: "elise",
		Age: 7,
	}
	fmt.Println(elise.EstMajeur())
}

/*
// Méthode 1 : Avec noms de champs (recommandé)
p1 := Personne{
	Nom: "Alice",
	Age: 25,
}

// Méthode 2 : Par ordre (à éviter)
p2 := Personne{"Bob", 30}

// Méthode 3 : Partielle
p3 := Personne{Nom: "Charlie"}  // Age = 0 par défaut

// Méthode 4 : Avec new (retourne un pointeur)
p4 := new(Personne)
p4.Nom = "David"
*/

/*

func main() {
	// Test avec majeur
	adulte := Personne{Nom: "Marie", Age: 25}
	fmt.Printf("%s (%d ans) - Majeur: %v\n", 
		adulte.Nom, adulte.Age, adulte.EstMajeur())
	
	// Test avec mineur
	enfant := Personne{Nom: "Élise", Age: 7}
	fmt.Printf("%s (%d ans) - Majeur: %v\n", 
		enfant.Nom, enfant.Age, enfant.EstMajeur())
	
	// Test cas limite
	limite := Personne{Nom: "Jean", Age: 18}
	fmt.Printf("%s (%d ans) - Majeur: %v\n", 
		limite.Nom, limite.Age, limite.EstMajeur())
}

// Résultat :
// Marie (25 ans) - Majeur: true
// Élise (7 ans) - Majeur: false
// Jean (18 ans) - Majeur: true
*/