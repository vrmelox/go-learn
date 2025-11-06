package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		ligne := scanner.Text()
		if len(strings.TrimSpace(ligne)) > 0 {
			ln++
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("erreur de lecture: %w", err)
	}
	return ln, nil
}

func main() {
	f, err := compterLignes("../danse.txt")

	if err != nil {
		fmt.Println("Ah oui, un problème occured:", err)
	}
	fmt.Println("Le nombre de ligne est de :",f)	
}

/*
func compterLignesGros(chemin string) (int, error) {
	f, err := os.Open(chemin)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	
	const maxCapacity = 1024 * 1024 // 1MB buffer
	scanner := bufio.NewScanner(f)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	
	count := 0
	for scanner.Scan() {
		count++
	}
	
	return count, scanner.Err()
}
	*/

	/*
func rechercherDansFichier(chemin, motif string) ([]string, error) {
	f, err := os.Open(chemin)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	
	var lignesTrouvees []string
	scanner := bufio.NewScanner(f)
	numLigne := 0
	
	for scanner.Scan() {
		numLigne++
		ligne := scanner.Text()
		if strings.Contains(ligne, motif) {
			lignesTrouvees = append(lignesTrouvees, 
				fmt.Sprintf("Ligne %d: %s", numLigne, ligne))
		}
	}
	
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	
	return lignesTrouvees, nil
}

func main() {
	resultats, err := rechercherDansFichier("../danse.txt", "danse")
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	
	fmt.Printf("🔍 %d ligne(s) trouvée(s):\n", len(resultats))
	for _, ligne := range resultats {
		fmt.Println(ligne)
	}
}