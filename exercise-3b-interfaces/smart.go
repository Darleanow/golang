package main

import "fmt"

// User holds the possible contact channels for a person.
type User struct {
	Name  string
	Email string
	Phone string
}

// processData reacts to the dynamic type behind an empty interface value.
func processData(data any) {
	switch v := data.(type) {
	case int:
		fmt.Printf("Donnée de type entier : %d\n", v)
	case string:
		fmt.Printf("Donnée de type chaîne : %s\n", v)
	case bool:
		fmt.Printf("Donnée de type booléen : %t\n", v)
	case *EmailNotifier:
		fmt.Printf("Donnée de type EmailNotifier pour %s\n", v.Recipient)
	default:
		fmt.Printf("Type de donnée inconnu : %T\n", v)
	}
}

// sendSmartNotification picks the right channel from the dynamic type of data
// and, for a User, from the contact information available.
func sendSmartNotification(data any, message string) error {
	switch v := data.(type) {
	case User:
		switch {
		case v.Email != "":
			return EmailNotifier{Recipient: v.Email, Sender: "system@app.io"}.Send(message)
		case v.Phone != "":
			return SMSNotifier{PhoneNumber: v.Phone}.Send(message)
		default:
			return ConsoleNotifier{}.Send("aucune méthode de contact pour " + v.Name)
		}
	case string:
		return ConsoleNotifier{}.Send("Message générique : " + message)
	default:
		return fmt.Errorf("type non supporté : %T", v)
	}
}
