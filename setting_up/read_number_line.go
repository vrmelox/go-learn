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
		return []string{}, errors.New("Le fichier " + fichier + " ne s'ouvre pas à cause de " + err.Error())
	}
	text := []string{}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ligne := scanner.Text()
		if len(strings.TrimSpace(ligne)) == 0 {
			return []string{}, errors.New("le fichier " + fichier + " contient des lignes vides")
		}
		text = append(text, ligne)
	}
	if err := scanner.Err(); err != nil {
		return []string{}, fmt.Errorf("erreur de lecture: %w", err)
	}
	return text, nil
}

func readNumberOnLine(fichier string) (int, error) {
	text, err := getContent(fichier)

	if (err != nil) {
		return 0, errors.New("une erreur semble survenue à la lecture du fichier")
	}
	a, err := strconv.Atoi(text[0])
	if err != nil {
		return -1, fmt.Errorf("houston, nous avons un problème: %w", err)
	}
	return a, nil
}