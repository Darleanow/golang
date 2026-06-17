package main

import "slices"

// categorieExiste reports whether nom is present in categories.
func categorieExiste(nom string, categories []string) bool {
	return slices.Contains(categories, nom)
}

// supprimerCategorie removes nom from categories, returning the slice unchanged
// when nom is absent.
func supprimerCategorie(nom string, categories []string) []string {
	i := slices.Index(categories, nom)
	if i == -1 {
		return categories
	}
	return slices.Delete(categories, i, i+1)
}
