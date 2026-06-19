// Command exercise-5c-json demonstrates JSON serialisation and deserialisation
// in Go: struct tags, omitempty, ignored fields, error handling, and custom
// marshalling.
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// json:"-" on MotDePasse prevents credentials from ever reaching the wire,
// regardless of how the value is set - safer than relying on callers to clear it.
type Personne struct {
	Nom        string `json:"full_name"`
	Age        int    `json:"age_in_years"`
	Email      string `json:"contact_email,omitempty"`
	Actif      bool   `json:"is_active"`
	MotDePasse string `json:"-"`
}

func exercice1et2() {
	fmt.Println("Exercice 1 & 2 - Sérialisation et struct tags")

	alice := Personne{Nom: "Alice Dupont", Age: 30, Email: "alice@example.com", Actif: true, MotDePasse: "secret"}
	bob := Personne{Nom: "Bob Martin", Age: 25, Actif: false}

	for _, p := range []Personne{alice, bob} {
		data, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			fmt.Printf("Erreur Marshal : %v\n", err)
			continue
		}
		fmt.Println(string(data))
	}
}

type Produit struct {
	ID      int     `json:"product_id"`
	Nom     string  `json:"item_name"`
	Prix    float64 `json:"unit_price"`
	EnStock bool    `json:"in_stock"`
}

func exercice3() {
	fmt.Println("\nExercice 3 - Désérialisation")

	jsonString := `{
		"product_id": 101,
		"item_name": "Clavier Mécanique",
		"unit_price": 79.99,
		"in_stock": true
	}`

	var p Produit
	if err := json.Unmarshal([]byte(jsonString), &p); err != nil {
		fmt.Printf("Erreur Unmarshal : %v\n", err)
		return
	}
	fmt.Printf("ID=%d  Nom=%s  Prix=%.2f  EnStock=%v\n", p.ID, p.Nom, p.Prix, p.EnStock)
}

func exercice4() {
	fmt.Println("\nExercice 4 - Gestion des erreurs")

	malformedJSON := `{
		"product_id": 102,
		"item_name": "Souris Gaming",
		"unit_price": 49.99,
		"in_stock": true,
	`

	wrongTypeJSON := `{
		"product_id": "103",
		"item_name": "Écran UltraWide",
		"unit_price": 399.99,
		"in_stock": true
	}`

	for label, s := range map[string]string{
		"JSON malformé":             malformedJSON,
		"Type de données incorrect": wrongTypeJSON,
	} {
		var p Produit
		if err := json.Unmarshal([]byte(s), &p); err != nil {
			fmt.Printf("  [%s] Erreur : %v\n", label, err)
		}
	}
}

type Livre struct {
	ID         int      `json:"book_id"`
	Titre      string   `json:"title"`
	Auteur     string   `json:"author_name"`
	Annee      int      `json:"publication_year"`
	Genres     []string `json:"genres,omitempty"`
	ISBN       string   `json:"isbn_code,omitempty"`
	Disponible bool     `json:"is_available"`
	DateAjout  UnixTime `json:"date_ajout"`
}

// UnixTime stores dates as Unix timestamps (int64) in JSON instead of the
// default RFC 3339 string - useful when interoperating with APIs or front-ends
// that expect numeric timestamps.
type UnixTime struct{ time.Time }

func (u UnixTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Unix())
}

func (u *UnixTime) UnmarshalJSON(data []byte) error {
	var ts int64
	if err := json.Unmarshal(data, &ts); err != nil {
		return err
	}
	u.Time = time.Unix(ts, 0)
	return nil
}

func exercice5() {
	fmt.Println("\nExercice 5 - Scénario complet (Livre)")

	complet := Livre{
		ID:         1,
		Titre:      "Le Guide du voyageur galactique",
		Auteur:     "Douglas Adams",
		Annee:      1979,
		Genres:     []string{"Science-fiction", "Humour"},
		ISBN:       "978-2-07-036822-8",
		Disponible: true,
		DateAjout:  UnixTime{time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)},
	}

	partiel := Livre{
		ID:         2,
		Titre:      "L'art de la guerre",
		Auteur:     "Sun Tzu",
		Annee:      -500,
		Disponible: false,
		DateAjout:  UnixTime{time.Now()},
	}

	fmt.Println("\nPartie A - Sérialisation")
	for _, l := range []Livre{complet, partiel} {
		data, err := json.MarshalIndent(l, "", "  ")
		if err != nil {
			fmt.Printf("Erreur : %v\n", err)
			continue
		}
		fmt.Println(string(data))
	}

	fmt.Println("\nPartie B - Désérialisation")
	data, err := json.Marshal(complet)
	if err != nil {
		fmt.Printf("Erreur Marshal : %v\n", err)
		return
	}
	var lu Livre
	if err := json.Unmarshal(data, &lu); err != nil {
		fmt.Printf("Erreur Unmarshal : %v\n", err)
		return
	}
	fmt.Printf("Désérialisé → ID=%d Titre=%q Auteur=%q Année=%d Genres=%v ISBN=%q Disponible=%v DateAjout=%s\n",
		lu.ID, lu.Titre, lu.Auteur, lu.Annee, lu.Genres, lu.ISBN, lu.Disponible,
		lu.DateAjout.UTC().Format(time.DateOnly))
}

func main() {
	exercice1et2()
	exercice3()
	exercice4()
	exercice5()
}
