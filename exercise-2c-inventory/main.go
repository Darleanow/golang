// Command exercise-2c-inventory manipulates slices and maps through a small inventory.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func partie1Categories() {
	fmt.Println("Partie 1 - Catégories (slices)")

	categories := []string{"Électronique", "Vêtements", "Livres"}
	categories = append(categories, "Alimentation", "Jouets")
	fmt.Println("Catégories :", categories)

	fmt.Println(`"Livres" existe ?`, categorieExiste("Livres", categories))
	fmt.Println(`"Meubles" existe ?`, categorieExiste("Meubles", categories))

	categories = supprimerCategorie("Vêtements", categories)
	fmt.Println(`Après suppression de "Vêtements" :`, categories)

	categories = supprimerCategorie("Meubles", categories)
	fmt.Println(`Après tentative de suppression de "Meubles" (absent) :`, categories)

	// len counts stored elements; cap is the size of the backing array, which
	// append grows on demand while delete only shrinks len.
	fmt.Printf("len = %d, cap = %d\n", len(categories), cap(categories))
}

func partie2Produits() (map[int]Produit, map[int]int) {
	fmt.Println("\nPartie 2 - Produits et stock (maps)")

	inventaire := make(map[int]Produit)
	stock := make(map[int]int)

	produits := []Produit{
		{ID: 1, Nom: "Casque Audio", Prix: 79.90, Categorie: "Électronique"},
		{ID: 2, Nom: "T-shirt", Prix: 19.99, Categorie: "Vêtements"},
		{ID: 3, Nom: "Roman", Prix: 12.50, Categorie: "Livres"},
	}
	stockInitial := map[int]int{1: 10, 2: 50, 3: 30}
	for _, p := range produits {
		inventaire[p.ID] = p
		stock[p.ID] = stockInitial[p.ID]
	}

	// A map value is not addressable, so it is copied, updated and stored back.
	casque := inventaire[1]
	casque.Prix = 69.90
	inventaire[1] = casque
	stock[2] = 45

	fmt.Println("Inventaire complet :")
	for id, p := range inventaire {
		afficherProduit(p, stock[id])
	}

	if p, q, ok := obtenirProduit(3, inventaire, stock); ok {
		fmt.Printf("Produit trouvé : %s, stock %d\n", p.Nom, q)
	}
	if _, _, ok := obtenirProduit(99, inventaire, stock); !ok {
		fmt.Println("Produit 99 introuvable")
	}

	delete(inventaire, 3)
	delete(stock, 3)
	if _, _, ok := obtenirProduit(3, inventaire, stock); !ok {
		fmt.Println("Produit 3 supprimé")
	}

	fmt.Printf("Stock du Casque avant la vente : %d\n", stock[1])

	if vendreProduit(1, 3, stock) {
		fmt.Printf("Vente de 3 unités, stock maintenant : %d\n", stock[1])
	}
	if !vendreProduit(1, 1000, stock) {
		fmt.Println("Impossible de vendre 1000 : stock insuffisant")
	}
	reapprovisionnerProduit(1, 20, stock)
	fmt.Printf("Stock du Casque après réapprovisionnement : %d\n", stock[1])

	return inventaire, stock
}

func partie3Combinaison(inventaire map[int]Produit, stock map[int]int) {
	fmt.Println("\nPartie 3a - Tri et index par catégorie")

	parCategorie := indexerParCategorie(inventaire)
	listerProduitsParCategorie("Électronique", inventaire, parCategorie)
	listerProduitsParCategorie("Vêtements", inventaire, parCategorie)
	listerProduitsParCategorie("Introuvable", inventaire, parCategorie)

	fmt.Println("Tri par prix :")
	for _, p := range trierParPrix(inventaire, true) {
		fmt.Printf("  %s : %.2f\n", p.Nom, p.Prix)
	}

	valeurStock := valeurDuStock("Électronique", inventaire, stock, parCategorie)
	fmt.Printf("Valeur du stock Électronique : %.2f\n", valeurStock)
}

func partie3Performance() {
	fmt.Println("\nPartie 3b - Performance des maps")

	const totalProduits = 100_000
	familles := []string{"Électronique", "Vêtements", "Livres", "Alimentation", "Jouets"}

	produitAleatoire := func(i int) Produit {
		return Produit{
			ID:        i,
			Nom:       fmt.Sprintf("Réf-%d", i),
			Prix:      rand.Float64() * 100,
			Categorie: familles[rand.Intn(len(familles))],
		}
	}

	start := time.Now()
	sansCapacite := map[int]Produit{}
	for i := 0; i < totalProduits; i++ {
		sansCapacite[i] = produitAleatoire(i)
	}
	tempsSansCapacite := time.Since(start)

	start = time.Now()
	avecCapacite := make(map[int]Produit, totalProduits)
	for i := 0; i < totalProduits; i++ {
		avecCapacite[i] = produitAleatoire(i)
	}
	tempsAvecCapacite := time.Since(start)

	fmt.Printf("Insertion de %d éléments sans make(cap) : %v\n", totalProduits, tempsSansCapacite)
	fmt.Printf("Insertion de %d éléments avec make(cap) : %v\n", totalProduits, tempsAvecCapacite)

	start = time.Now()
	for i := 0; i < 10_000; i++ {
		_ = avecCapacite[rand.Intn(totalProduits)]
	}
	fmt.Printf("10 000 lectures aléatoires : %v\n", time.Since(start))

	start = time.Now()
	var valeurTotale float64
	for _, p := range avecCapacite {
		valeurTotale += p.Prix
	}
	fmt.Printf("Parcours complet : %v\n", time.Since(start))
}

func main() {
	partie1Categories()
	inventaire, stock := partie2Produits()
	partie3Combinaison(inventaire, stock)
	partie3Performance()
}
