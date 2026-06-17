package main

import "fmt"

// Produit is a single inventory item.
type Produit struct {
	ID        int
	Nom       string
	Prix      float64
	Categorie string
}

// obtenirProduit returns the product, its stock and whether it exists.
func obtenirProduit(id int, inventaire map[int]Produit, stock map[int]int) (Produit, int, bool) {
	p, ok := inventaire[id]
	if !ok {
		return Produit{}, 0, false
	}
	return p, stock[id], true
}

// vendreProduit decrements the stock when enough units are available and
// reports whether the sale succeeded.
func vendreProduit(id, quantite int, stock map[int]int) bool {
	disponible, ok := stock[id]
	if !ok || disponible < quantite {
		return false
	}
	stock[id] = disponible - quantite
	return true
}

// reapprovisionnerProduit increments the stock of a product.
func reapprovisionnerProduit(id, quantite int, stock map[int]int) {
	stock[id] += quantite
}

// afficherProduit prints one product line with its current stock.
func afficherProduit(p Produit, quantite int) {
	fmt.Printf("ID : %d | %-12s | %7.2f EUR | %-14s | stock %d\n",
		p.ID, p.Nom, p.Prix, p.Categorie, quantite)
}
