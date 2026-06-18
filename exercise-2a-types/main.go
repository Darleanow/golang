// Command exercise-2a-types explores types, variables and constants.
package main

import "fmt"

func explicitDeclarations() {
	fmt.Println("Exercice 1 - Déclarations explicites")

	var userName string = "Darleanow"
	var userAge int = 23
	var isLoggedIn bool = true
	var accountBalance float64 = 1500.75

	fmt.Println("Nom      :", userName)
	fmt.Println("Âge      :", userAge)
	fmt.Println("Connecté :", isLoggedIn)
	fmt.Println("Solde    :", accountBalance)
}

func typeInference() {
	fmt.Println("\nExercice 2 - Inférence de type (:=)")

	city := "Rennes"
	postalCode := 35000
	discountRate := 0.2

	fmt.Printf("ville       = %v (type %T)\n", city, city)
	fmt.Printf("code postal = %v (type %T)\n", postalCode, postalCode)
	fmt.Printf("remise      = %v (type %T)\n", discountRate, discountRate)
}

func constants() {
	fmt.Println("\nExercice 3 - Constantes")

	const pi = 3.14159
	const appName = "Gestionnaire Go"
	const releaseYear = 2023

	radius := 10.5
	fmt.Printf("Circonférence (rayon %.1f) = %.4f\n", radius, 2*pi*radius)
	fmt.Println("Appli :", appName)
	fmt.Println("Année :", releaseYear)

	// releaseYear = 2024 // won't compile: a constant can't be reassigned
}

func reassignmentAndZeroValues() {
	fmt.Println("\nExercice 4 - Réaffectation et valeurs par défaut")

	userAge := 23
	fmt.Println("Âge avant anniversaire :", userAge)
	userAge = 24
	fmt.Println("Âge après anniversaire :", userAge)

	// No initializer => zero value.
	var message string
	var counter int
	fmt.Printf("message vaut %q, compteur vaut %d\n", message, counter)
}

func bonus() {
	fmt.Println("\nBonus")

	var a, b, c int = 1, 2, 3
	fmt.Println("Déclaration multiple :", a, b, c)

	type Weekday int
	const (
		Monday Weekday = iota
		Tuesday
		Wednesday
	)
	fmt.Println("iota :", Monday, Tuesday, Wednesday)

	// int and float64 can't be mixed without an explicit conversion.
	whole := 10
	fraction := 2.5
	fmt.Printf("float64(%d) + %.1f = %.1f\n", whole, fraction, float64(whole)+fraction)
}

func main() {
	explicitDeclarations()
	typeInference()
	constants()
	reassignmentAndZeroValues()
	bonus()
}
