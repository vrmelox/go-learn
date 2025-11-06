package main

import (
	"fmt"
	"math"
)

// EXERCICE 9 : Interfaces
// Créez une interface Forme avec méthode Aire()
// Implémentez-la pour Rectangle et Cercle

type Forme interface {
	Aire() float64
}

type Rectangle struct {
	Longueur, Largeur float64
}

func (rectou Rectangle) Aire() float64 {
	return rectou.Longueur * rectou.Largeur
}

type Cercle struct {
	Rayon float64
}

func (cer Cercle) Aire() float64 {
	return math.Pi * math.Pow(cer.Rayon, 2.0)
}
// Implémentez les méthodes Aire() pour les deux structures

func main() {
	rec := Rectangle{Longueur:4.5, Largeur:6.5}
	cer := Cercle{Rayon: 5.2}
	fmt.Println("L'aire du rectangle est:", Forme.Aire(rec))
	fmt.Println("L'aire du cercle est:", Forme.Aire(cer))
}

/*
====================POLYMORPHISME==================
// Fonction qui fonctionne avec TOUTES les formes
func calculerAireTotal(formes []Forme) float64 {
	total := 0.0
	for _, forme := range formes {
		total += forme.Aire()
	}
	return total
}

func main() {
	formes := []Forme{
		Rectangle{Longueur: 4.5, Largeur: 6.5},
		Cercle{Rayon: 5.2},
		Rectangle{Longueur: 10, Largeur: 20},
		Cercle{Rayon: 3},
	}
	
	total := calculerAireTotal(formes)
	fmt.Printf("Aire totale: %.2f\n", total)
	
	// Afficher chaque forme
	for i, forme := range formes {
		fmt.Printf("Forme %d - Aire: %.2f\n", i+1, forme.Aire())
	}
}
=====================INTERFACES AVEC PLUSIEURS METHODES===================
type Forme interface {
	Aire() float64
	Perimetre() float64
}

type Rectangle struct {
	Longueur, Largeur float64
}

func (r Rectangle) Aire() float64 {
	return r.Longueur * r.Largeur
}

func (r Rectangle) Perimetre() float64 {
	return 2 * (r.Longueur + r.Largeur)
}

type Cercle struct {
	Rayon float64
}

func (c Cercle) Aire() float64 {
	return math.Pi * c.Rayon * c.Rayon
}

func (c Cercle) Perimetre() float64 {
	return 2 * math.Pi * c.Rayon
}
==========================Type assertion (vérifier le type réel)===========
func decrireForme(f Forme) {
	fmt.Printf("Aire: %.2f\n", f.Aire())
	
	// Type assertion
	switch forme := f.(type) {
	case Rectangle:
		fmt.Printf("Rectangle %v x %v\n", forme.Longueur, forme.Largeur)
	case Cercle:
		fmt.Printf("Cercle de rayon %v\n", forme.Rayon)
	default:
		fmt.Println("Forme inconnue")
	}
}

// Ou avec vérification simple
func main() {
	var f Forme = Cercle{Rayon: 5}
	
	// Vérifier si c'est un Cercle
	if cercle, ok := f.(Cercle); ok {
		fmt.Printf("C'est un cercle de rayon %.2f\n", cercle.Rayon)
	}
}
*/