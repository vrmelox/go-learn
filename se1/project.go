package main

import ("fmt";"slices")

/*
=============================================================================
PROJET 1 : CALCULATRICE DE STATISTIQUES
=============================================================================
Créez un programme qui analyse une liste de notes d'étudiants et calcule :
- La moyenne
- La note minimale et maximale
- Le nombre de notes >= 10 (réussite)
- Le taux de réussite en pourcentage

Utilisez toutes les compétences des exercices 1-5
*/

func sumSlice(notes []float64) float64 {
	if len(notes) == 0 {
		return 0
	}
	total := 0.0
	for _, note := range notes {
		total += note
	}
	return total
}

func isTen(notes []float64) int {
	count := 0
	for _, elem := range notes {
		if elem >= 10 {
			count++
		}
	}
	return count
}

type StatistiquesNotes struct {
	Moyenne      float64
	Min          float64
	Max          float64
	NbReussites  int
	TauxReussite float64
}

func analyserNotes(notes []float64) StatistiquesNotes {
	statis := StatistiquesNotes {}
	statis.Moyenne = sumSlice(notes)/float64(len(notes))
	statis.Min = slices.Min(notes)
	statis.Max = slices.Max(notes)
	statis.NbReussites = isTen(notes)
	statis.TauxReussite = float64(statis.NbReussites) * 100.0 / float64(len(notes))
	return statis
}

func main() {
	notes := []float64{12.5, 15.3, 8.7, 19.2, 10.4,14.8, 7.6, 13.9, 9.1, 16.7,11.2, 18.4, 6.5, 17.9, 13.3,10.0, 8.1, 15.7, 19.8, 9.9,14.0, 7.2, 12.8, 16.1, 11.7,}
	za := analyserNotes(notes)
	fmt.Println(za)
}