package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
	"strconv"
)

func getContent(fichier string) ([]string, error) {
	f, err := os.Open(fichier)
	if (err != nil) {
		return []string{}, fmt.Errorf("Le fichier %w ne s'ouvre pas à cause de %w", fichier, err)
	}
	text := []string{}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ligne := scanner.Text()
		if len(strings.TrimSpace(ligne)) == 0 {
			return []string{}, fmt.Errorf("Le fichier %w contient des lignes vides", fichier)
		}
		text = append(text, ligne + "\n")
	}
	if err := scanner.Err(); err != nil {
		return []string{}, fmt.Errorf("erreur de lecture: %w", err)
	}
	return text, nil
}

func readNumberOnLine(fichier string) (int, error) {
	text, err := getContent(fichier)

	if (err != nil) {
		return 0, errors.New("Une erreur semble survenue à la lecture du fichier")
	}
	a, err := strconv.Atoi(text[0])
	if err != nil {
		return -1, fmt.Errorf("Huston, nous avons un problème: %w", err)
	}
	return a, nil
}

func main() {
	file := os.Args[1]
	nbr, err := readNumberOnLine(file)
	if err != nil {
		fmt.Errorf("Fichier invalide")
	}
	fmt.Printf("Le nombre est de %d\n", nbr)
}