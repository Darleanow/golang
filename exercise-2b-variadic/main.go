// Command exercise-2b-variadic implements variadic functions and multiple return values.
package main

import (
	"errors"
	"fmt"
	"math"
)

// CalculerStatistiquesBase returns the sum, count and average of numbers.
func CalculerStatistiquesBase(numbers ...int) (sum, count int, average float64) {
	count = len(numbers)
	if count == 0 {
		return 0, 0, 0
	}
	for _, n := range numbers {
		sum += n
	}
	return sum, count, float64(sum) / float64(count)
}

// CalculerStatistiquesCompletes returns min, max, sum, average and count.
// It fails when called without arguments.
func CalculerStatistiquesCompletes(numbers ...float64) (minVal, maxVal, sum, average float64, count int, err error) {
	count = len(numbers)
	if count == 0 {
		return 0, 0, 0, 0, 0, errors.New("aucun argument fourni")
	}

	minVal = math.MaxFloat64
	maxVal = -math.MaxFloat64
	for _, n := range numbers {
		sum += n
		minVal = min(minVal, n)
		maxVal = max(maxVal, n)
	}
	return minVal, maxVal, sum, sum / float64(count), count, nil
}

// AnalyserDonneesCapteur keeps the readings in ]0, 100] (Celsius), computes
// their stats and counts how many were rejected.
func AnalyserDonneesCapteur(readings ...float64) (minVal, maxVal, average float64, validCount, invalidCount int, err error) {
	valid := make([]float64, 0, len(readings))
	for _, r := range readings {
		if r > 0 && r <= 100 {
			valid = append(valid, r)
		} else {
			invalidCount++
		}
	}

	if len(valid) == 0 {
		return 0, 0, 0, 0, invalidCount, errors.New("aucun relevé valide trouvé")
	}

	minVal, maxVal, _, average, validCount, err = CalculerStatistiquesCompletes(valid...)
	return minVal, maxVal, average, validCount, invalidCount, err
}

func main() {
	fmt.Println("Exercice 1 — Statistiques de base")
	somme, nombre, moyenne := CalculerStatistiquesBase(10, 20, 30, 40)
	fmt.Printf("Somme : %d, Nombre : %d, Moyenne : %.2f\n", somme, nombre, moyenne)

	somme, nombre, moyenne = CalculerStatistiquesBase(42)
	fmt.Printf("Un seul : Somme : %d, Nombre : %d, Moyenne : %.2f\n", somme, nombre, moyenne)

	somme, nombre, moyenne = CalculerStatistiquesBase()
	fmt.Printf("Vide : Somme : %d, Nombre : %d, Moyenne : %.2f\n", somme, nombre, moyenne)

	fmt.Println("\nExercice 2 — Statistiques complètes")
	if lo, hi, total, mean, n, err := CalculerStatistiquesCompletes(1.5, 2.8, 0.7, 3.1); err != nil {
		fmt.Println("Erreur :", err)
	} else {
		fmt.Printf("Min : %.2f, Max : %.2f, Somme : %.2f, Moyenne : %.2f, Nombre : %d\n", lo, hi, total, mean, n)
	}
	if _, _, _, _, _, err := CalculerStatistiquesCompletes(); err != nil {
		fmt.Println("Entrée vide :", err)
	}

	fmt.Println("\nExercice 3 — Analyse de données de capteur")
	if lo, hi, mean, ok, ko, err := AnalyserDonneesCapteur(22.5, 23.1, -5.0, 101.0, 21.9, 0.0, 24.0); err != nil {
		fmt.Println("Erreur d'analyse :", err)
	} else {
		fmt.Printf("Temp Min : %.2f, Max : %.2f, Moyenne : %.2f, Valides : %d, Invalides : %d\n", lo, hi, mean, ok, ko)
	}
	if _, _, _, _, _, err := AnalyserDonneesCapteur(-10.0, 105.0, 0.0); err != nil {
		fmt.Println("Toutes invalides :", err)
	}
}
