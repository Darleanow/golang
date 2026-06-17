package main

import (
	"fmt"
	"sort"
)

// indexerParCategorie groups product IDs by their category.
func indexerParCategorie(inventaire map[int]Produit) map[string][]int {
	index := make(map[string][]int)
	for id, p := range inventaire {
		index[p.Categorie] = append(index[p.Categorie], id)
	}
	return index
}

// listerProduitsParCategorie prints every product of the given category.
func listerProduitsParCategorie(categorie string, inventaire map[int]Produit, parCategorie map[string][]int) {
	ids := parCategorie[categorie]
	if len(ids) == 0 {
		fmt.Printf("Aucun produit dans la catégorie %q\n", categorie)
		return
	}
	fmt.Printf("Produits dans la catégorie %q :\n", categorie)
	for _, id := range ids {
		p := inventaire[id]
		fmt.Printf("  - %s (%.2f EUR)\n", p.Nom, p.Prix)
	}
}

// trierParPrix returns the products sorted by price, ascending when ascendant
// is true and descending otherwise.
func trierParPrix(inventaire map[int]Produit, ascendant bool) []Produit {
	produits := make([]Produit, 0, len(inventaire))
	for _, p := range inventaire {
		produits = append(produits, p)
	}

	sort.Slice(produits, func(i, j int) bool {
		if ascendant {
			return produits[i].Prix < produits[j].Prix
		}
		return produits[i].Prix > produits[j].Prix
	})
	return produits
}

// valeurDuStock returns the total stock value of a category.
func valeurDuStock(categorie string, inventaire map[int]Produit, stock map[int]int, parCategorie map[string][]int) float64 {
	total := 0.0
	for _, id := range parCategorie[categorie] {
		total += inventaire[id].Prix * float64(stock[id])
	}
	return total
}
