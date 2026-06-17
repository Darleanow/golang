package main

import "fmt"

func exo1() {
	nomUtilisateur := "Darleanow"
	var ageUtilisateur uint8 = 23
	estConnecte := true
	var soldeCompte float64 = 99999999

	fmt.Println(
		nomUtilisateur,
		"a",
		ageUtilisateur,
		"ans, connecté :",
		estConnecte,
		"et a",
		soldeCompte,
		"sur son compte",
	)
}

func exo2() {
	villeResidence := "Secret"
	codePostal := 11111
	tauxRemise := 20.0

	fmt.Printf(
		"%s, %T, %d, %T, %f, %T\n",
		villeResidence,
		villeResidence,
		codePostal,
		codePostal,
		tauxRemise,
		tauxRemise,
	)
}

func exo3() {
	const PI = 3.14159
	const NOM_APPLICATION = "Gestionnaire Go"
	const ANNEE_LANCEMENT = 2023

	rayon := 10.5

	fmt.Println("La circonférence d'un cercle de rayon", rayon, "est", 2*rayon*PI)

	// ANNEE_LANCEMENT = 2024
	// ./main.go:48:2: cannot assign to ANNEE_LANCEMENT (neither addressable nor a map index expression)
}

func exo4() {
	var ageUtilisateur uint8 = 23

	fmt.Println("Age utilisateur:", ageUtilisateur)

	ageUtilisateur = 24
	fmt.Println("Age utilisateur après modification:", ageUtilisateur)

	var message string
	fmt.Println("Contenu de message:", message)

	var compteur int
	fmt.Println("Contenu de compteur:", compteur)
}

func bonus() {
	var a, b, c int

	fmt.Println(a, b, c)

	type Weekday int

	const (
		Monday Weekday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

	d := 1
	e := 2.5
	f := float64(d) + e
	fmt.Println(f)
}

func main() {
	exo1()
	exo2()
	exo3()
	exo4()
	bonus()
}
