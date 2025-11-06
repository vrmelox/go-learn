package main

import "fmt"

// EXERCICE 8 : Pointeurs
// Créez une fonction qui incrémente une valeur via pointeur
// Input: *int
// Output: rien (modifie directement)

func incrementer(n *int) {
	*n++
}

func main() {
	age := 18
	p := &age
	incrementer(p)
	fmt.Println(*p)
}