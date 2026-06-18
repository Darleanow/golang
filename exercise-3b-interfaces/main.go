// Command exercise-3b-interfaces implements a small notification system to
// practise Go interfaces: implicit satisfaction, polymorphism and the empty
// interface with type switches.
package main

import "fmt"

func main() {
	fmt.Println("Partie 1 - Notifications polymorphiques")
	notifiers := []Notifier{
		EmailNotifier{Recipient: "bob@mail.com", Sender: "alice@mail.com"},
		SMSNotifier{PhoneNumber: "+33612345678"},
		ConsoleNotifier{},
	}
	for _, n := range notifiers {
		if err := n.Send("Votre commande est prête !"); err != nil {
			fmt.Println("erreur d'envoi :", err)
		}
	}

	fmt.Println("\nPartie 2 - Interface vide")
	processData(42)
	processData("Bonjour le monde")
	processData(true)
	processData(&EmailNotifier{Recipient: "carol@mail.com", Sender: "system@app.io"})
	processData([]int{1, 2, 3})
	processData(3.14)

	fmt.Println("\nPartie 3 - Notification intelligente")
	cases := []any{
		User{Name: "Alice", Email: "alice@mail.com", Phone: "+33600000000"},
		User{Name: "Bob", Phone: "+33611111111"},
		User{Name: "Carol"},
		"info diffusée",
		99,
	}
	for _, data := range cases {
		if err := sendSmartNotification(data, "Notification importante"); err != nil {
			fmt.Println("erreur :", err)
		}
	}
}
