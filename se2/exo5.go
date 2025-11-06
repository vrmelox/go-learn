package main

import (
	"fmt"
	"errors"
)
// EXERCICE 10 : Méthodes avec pointeurs
// Créez une structure CompteBancaire avec solde
// Créez des méthodes Deposer() et Retirer() qui modifient le solde

type CompteBancaire struct {
	Solde float64
}

func (c *CompteBancaire) Deposer(montant float64) {
	if montant < 0 {
		errors.New("Le montant ne peut pas être négatif")
	}
	c.Solde += montant
}

func (c *CompteBancaire) Retirer(montant float64) error {
	if c.Solde < montant {
		return  errors.New("Solde insuffisant")
	} else if montant < 0 {
		return errors.New("Le montant ne peut pas être négatif")
	}
	c.Solde -= montant
	return nil
}

func main() {
	bank := CompteBancaire{Solde: 1452.0}
	bank.Retirer(-47.0)
	bank.Retirer(875)
	bank.Deposer(74552)
	fmt.Println(bank.Solde)
}

/*
// Version 1 : Avec historique des transactions
type Transaction struct {
	Type    string  // "Dépôt" ou "Retrait"
	Montant float64
	Date    string
}

type CompteBancaire struct {
	Solde        float64
	Historique   []Transaction
	NumeroCompte string
}

func (c *CompteBancaire) Deposer(montant float64) error {
	if montant <= 0 {
		return errors.New("montant invalide")
	}
	c.Solde += montant
	c.Historique = append(c.Historique, Transaction{
		Type:    "Dépôt",
		Montant: montant,
		Date:    time.Now().Format("2006-01-02 15:04:05"),
	})
	return nil
}

func (c *CompteBancaire) Retirer(montant float64) error {
	if montant <= 0 {
		return errors.New("montant invalide")
	}
	if c.Solde < montant {
		return fmt.Errorf("solde insuffisant")
	}
	c.Solde -= montant
	c.Historique = append(c.Historique, Transaction{
		Type:    "Retrait",
		Montant: montant,
		Date:    time.Now().Format("2006-01-02 15:04:05"),
	})
	return nil
}

func (c *CompteBancaire) AfficherHistorique() {
	fmt.Println("\n📜 Historique des transactions:")
	for i, t := range c.Historique {
		fmt.Printf("%d. %s: %.2f € [%s]\n", i+1, t.Type, t.Montant, t.Date)
	}
}

// Version 2 : Avec limite de découvert
type CompteBancaire struct {
	Solde             float64
	DecouvertAutorise float64
}

func (c *CompteBancaire) Retirer(montant float64) error {
	if montant <= 0 {
		return errors.New("montant invalide")
	}
	soldeApres := c.Solde - montant
	if soldeApres < -c.DecouvertAutorise {
		return fmt.Errorf("découvert dépassé: limite %.2f €", c.DecouvertAutorise)
	}
	c.Solde = soldeApres
	return nil
}

// Version 3 : Avec méthodes d'information
func (c *CompteBancaire) ConsulterSolde() float64 {
	return c.Solde
}

func (c *CompteBancaire) EstDansLeRouge() bool {
	return c.Solde < 0
}

func (c *CompteBancaire) Transferer(destination *CompteBancaire, montant float64) error {
	if err := c.Retirer(montant); err != nil {
		return fmt.Errorf("transfert échoué: %w", err)
	}
	if err := destination.Deposer(montant); err != nil {
		// Annuler le retrait si le dépôt échoue
		c.Deposer(montant)
		return fmt.Errorf("transfert échoué: %w", err)
	}
	return nil
}

// Version 4 : Avec constructeur
func NouveauCompteBancaire(soldeInitial float64, numero string) (*CompteBancaire, error) {
	if soldeInitial < 0 {
		return nil, errors.New("solde initial ne peut pas être négatif")
	}
	return &CompteBancaire{
		Solde:        soldeInitial,
		NumeroCompte: numero,
		Historique:   []Transaction{},
	}, nil
}*/