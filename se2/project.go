package main

import (
	"errors"
	"fmt"
)

/*
=============================================================================
PROJET 2 : GESTIONNAIRE DE BIBLIOTHÈQUE
=============================================================================
Créez un système de gestion de livres avec :
- Structure Livre (titre, auteur, annee, disponible)
- Structure Bibliotheque avec slice de livres
- Méthodes : AjouterLivre, EmprunterLivre, RetournerLivre
- Méthode RechercherParAuteur qui retourne tous les livres d'un auteur
- Méthode LivresDisponibles qui retourne le nombre de livres disponibles
*/

type Livre struct {
	Titre      string
	Auteur     string
	Annee      int
	Disponible bool
}

type Bibliotheque struct {
	Livres []Livre
}

func (bn Bibliotheque) EmprunterLivre(book Livre) Bibliotheque {
	bn.Livres = append(bn.Livres, book)
	return  bn
}
func CompareLivre(book1, book2 Livre) bool {
	if book1.Annee == book2.Annee && book1.Auteur == book2.Auteur && book1.Disponible == book2.Disponible && book1.Titre == book2.Titre {
		return true
	}
	return false
}
func (bn Bibliotheque) AjouterLivre(book Livre) Bibliotheque {
	for _, livre := range bn.Livres {
		if CompareLivre(livre, book) && livre.Disponible != false {
			livre.Disponible = false
			return bn
		}
	}
	return  bn
}

func (bn Bibliotheque) RetournerLivre(book Livre) Bibliotheque {
	for i := range bn.Livres {
		if CompareLivre(bn.Livres[i], book) && bn.Livres[i].Disponible != true {
			bn.Livres[i].Disponible = true
			return bn
		}
	}
	return  bn
}

func (bn Bibliotheque) RechercherParAuteur(author string) ([]Livre, error) {
	var bookis []Livre
	for _, livre := range bn.Livres {
		if (livre.Auteur == author) {
			bookis = append(bookis, livre)
		}
	}
	if len(bookis) == 0 {
		return bookis, errors.New("Aucun livre n'existe avec ce nom d'auteur")
	}
	return bookis, nil
}

func (bn Bibliotheque) LivresDisponibles() int {
	nbr := 0
	for _, livre := range bn.Livres {
		if livre.Disponible !=false {
			nbr++
		}
	}
	return nbr
}

// Implémentez toutes les méthodes nécessaires

/*
package main

import (
	"errors"
	"fmt"
	"strings"
)

type Livre struct {
	Titre      string
	Auteur     string
	Annee      int
	Disponible bool
}

type Bibliotheque struct {
	Livres []Livre
}

// Ajouter un nouveau livre à la bibliothèque
func (b *Bibliotheque) AjouterLivre(livre Livre) {
	livre.Disponible = true  // Nouveau livre = disponible
	b.Livres = append(b.Livres, livre)
}

// Emprunter un livre (le rendre indisponible)
func (b *Bibliotheque) EmprunterLivre(titre string) error {
	for i := range b.Livres {
		if b.Livres[i].Titre == titre {
			if !b.Livres[i].Disponible {
				return errors.New("livre déjà emprunté")
			}
			b.Livres[i].Disponible = false
			return nil
		}
	}
	return errors.New("livre non trouvé")
}

// Retourner un livre (le rendre disponible)
func (b *Bibliotheque) RetournerLivre(titre string) error {
	for i := range b.Livres {
		if b.Livres[i].Titre == titre {
			if b.Livres[i].Disponible {
				return errors.New("livre déjà disponible")
			}
			b.Livres[i].Disponible = true
			return nil
		}
	}
	return errors.New("livre non trouvé")
}

// Rechercher tous les livres d'un auteur
func (b *Bibliotheque) RechercherParAuteur(auteur string) []Livre {
	var resultat []Livre
	auteurLower := strings.ToLower(auteur)
	
	for _, livre := range b.Livres {
		if strings.ToLower(livre.Auteur) == auteurLower {
			resultat = append(resultat, livre)
		}
	}
	return resultat
}

// Compter les livres disponibles
func (b *Bibliotheque) LivresDisponibles() int {
	count := 0
	for _, livre := range b.Livres {
		if livre.Disponible {
			count++
		}
	}
	return count
}

// Afficher tous les livres
func (b *Bibliotheque) AfficherCatalogue() {
	fmt.Println("\n📚 CATALOGUE DE LA BIBLIOTHÈQUE")
	fmt.Println(strings.Repeat("=", 60))
	for i, livre := range b.Livres {
		statut := "✅ Disponible"
		if !livre.Disponible {
			statut = "❌ Emprunté"
		}
		fmt.Printf("%d. %s\n", i+1, livre.Titre)
		fmt.Printf("   Auteur: %s | Année: %d | %s\n", 
			livre.Auteur, livre.Annee, statut)
	}
	fmt.Printf("\nTotal: %d livres | Disponibles: %d\n", 
		len(b.Livres), b.LivresDisponibles())
}

func main() {
	// Créer une bibliothèque
	biblio := &Bibliotheque{}
	
	// Ajouter des livres
	biblio.AjouterLivre(Livre{
		Titre:  "1984",
		Auteur: "George Orwell",
		Annee:  1949,
	})
	biblio.AjouterLivre(Livre{
		Titre:  "La Ferme des animaux",
		Auteur: "George Orwell",
		Annee:  1945,
	})
	biblio.AjouterLivre(Livre{
		Titre:  "Le Seigneur des Anneaux",
		Auteur: "J.R.R. Tolkien",
		Annee:  1954,
	})
	biblio.AjouterLivre(Livre{
		Titre:  "Le Hobbit",
		Auteur: "J.R.R. Tolkien",
		Annee:  1937,
	})
	
	// Afficher le catalogue
	biblio.AfficherCatalogue()
	
	// Emprunter un livre
	fmt.Println("\n📖 Emprunt de '1984'...")
	if err := biblio.EmprunterLivre("1984"); err != nil {
		fmt.Println("Erreur:", err)
	} else {
		fmt.Println("✅ Livre emprunté avec succès")
	}
	
	// Rechercher par auteur
	fmt.Println("\n🔍 Livres de George Orwell:")
	livresOrwell := biblio.RechercherParAuteur("George Orwell")
	for _, livre := range livresOrwell {
		statut := "disponible"
		if !livre.Disponible {
			statut = "emprunté"
		}
		fmt.Printf("  - %s (%d) [%s]\n", livre.Titre, livre.Annee, statut)
	}
	// Statistiques
	fmt.Printf("\n📊 Livres disponibles: %d/%d\n", 
		biblio.LivresDisponibles(), len(biblio.Livres))
	// Retourner le livre
	fmt.Println("\n📥 Retour de '1984'...")
	if err := biblio.RetournerLivre("1984"); err != nil {
		fmt.Println("Erreur:", err)
	} else {
		fmt.Println("✅ Livre retourné avec succès")
	}
	// Catalogue final
	biblio.AfficherCatalogue()
}

*/