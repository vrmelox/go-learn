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
		return errors.New("Toutes les données n'ont pas pu être écrites")
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

/*func ajouterLignes(chemin string, lignes []string) error {
	// O_APPEND pour ajouter, O_CREATE pour créer si n'existe pas
	f, err := os.OpenFile(chemin, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("erreur d'ouverture: %w", err)
	}
	defer f.Close()
	
	writer := bufio.NewWriter(f)
	defer writer.Flush()
	
	for _, ligne := range lignes {
		if _, err := writer.WriteString(ligne + "\n"); err != nil {
			return err
		}
	}
	
	return nil
}

/*
type ResultatEcriture struct {
	NbLignes      int
	NbCaracteres  int
	TailleFichier int64
}

func ecrireLignesAvecStats(chemin string, lignes []string) (ResultatEcriture, error) {
	f, err := os.Create(chemin)
	if err != nil {
		return ResultatEcriture{}, err
	}
	defer f.Close()
	
	stats := ResultatEcriture{}
	writer := bufio.NewWriter(f)
	defer writer.Flush()
	
	for _, ligne := range lignes {
		n, err := writer.WriteString(ligne + "\n")
		if err != nil {
			return stats, err
		}
		stats.NbLignes++
		stats.NbCaracteres += n
	}
	
	// Obtenir la taille du fichier
	info, err := f.Stat()
	if err == nil {
		stats.TailleFichier = info.Size()
	}
	
	return stats, nil
}

func main() {
	data := []string{"Ligne 1", "Ligne 2", "Ligne 3"}
	stats, err := ecrireLignesAvecStats("output.txt", data)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	
	fmt.Printf("📊 Statistiques d'écriture:\n")
	fmt.Printf("  Lignes écrites: %d\n", stats.NbLignes)
	fmt.Printf("  Caractères: %d\n", stats.NbCaracteres)
	fmt.Printf("  Taille du fichier: %d octets\n", stats.TailleFichier)
}
/*
func copierFichier(source, destination string) error {
	// Lire le fichier source
	lignes, err := lireToutesLesLignes(source)
	if err != nil {
		return fmt.Errorf("erreur de lecture: %w", err)
	}
	
	// Écrire dans la destination
	if err := ecrireLignes(destination, lignes); err != nil {
		return fmt.Errorf("erreur d'écriture: %w", err)
	}
	
	return nil
}

func lireToutesLesLignes(chemin string) ([]string, error) {
	f, err := os.Open(chemin)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	
	var lignes []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lignes = append(lignes, scanner.Text())
	}
	
	return lignes, scanner.Err()
}