package main

import "fmt"

type basket struct {
	name    string
	color   string
	content []string
}

type clothing struct {
	name  string
	color string
	state string
}

func main() {
	Vetement1 := clothing{name: "tshirt", color: "blanc", state: "propre"}
	Vetement2 := clothing{name: "tshirt", color: "noir", state: "propre"}
	Vetement3 := clothing{name: "tshirt", color: "couleur", state: "propre"}

	Corbeille1 := backet{name: "pannier_blanc", color: "blanc", content: Vetement1}
	Corbeille2 := backet{name: "pannier_noir", color: "noir", content: Vetement2}
	Corbeille3 := backet{name: "pannier_couleur", color: "couleur", content: Vetement3}
}

func (b basket) displayBasket() {
	if len(b.content) < 1 {
		fmt.Printf("===Votre panier est vide===")
	} else {
		for _, vetements := range b.content {
			fmt.Println("=== panier %s ===", b.name)
			fmt.Printf("- t-shirt %s", c.color)
			fmt.Printf("\t couleur : %s", c.color)
			fmt.Printf("\t etat : %s", c.state)
		}
	}
	return fmt.Println()
}

func (b basket) addToBasket(vetement clothing) {
	if vetement.color != b.color {
		fmt.Println("Ajout impossible, la couleur du vetement ne correspond pas !")
	} else {
		content += append(vetement, content)
		fmt.Println("Ajout d'un nouveau vetement : t-shirt %s au panier %s", vetement.color, b.color)
	}
	return fmt.Println()
}

func (b basket) cleanBasket(compteur *int, vetement clothing) {
	if len(b.content) < 1 {
		fmt.Println("Il n'y a pas de vetement dans le panier a linge")
	} else {
		for _, vetement := range b.content {
			&compteur++
			vetement.state = "propre"
			fmt.Printf("=== Nettoyage du panier couleur ===")
			fmt.Printf("Le lavage du panier %s est termine, %d vetement ont ete laves", b.color, compteur)
		}
	}
	return fmt.Println()
}

func (b basket) emptyCleanLaundry(vetement clothing) {
	if b.content != "" {
		Println("=== Vidage du panier en cours ===")
		Println("Les vetements suivants ont ete retires")
		for _, vetement := range b.content {
			if vetement.state == "propre" {
				b.content := append(b.content, "")
				fmt.Printf("- t-shirt %s", vetement.color)
			}
		}
	} else {
		return fmt.Println("Votre panier est deja vide !")
	}
}
