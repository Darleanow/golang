package main

import "fmt"

// Notifier is the behaviour shared by every notification channel. Any type with
// a Send(string) error method satisfies it implicitly.
type Notifier interface {
	Send(message string) error
}

// EmailNotifier sends notifications by e-mail.
type EmailNotifier struct {
	Recipient string
	Sender    string
}

func (e EmailNotifier) Send(message string) error {
	fmt.Printf("[EMAIL] De %s à %s : %s\n", e.Sender, e.Recipient, message)
	return nil
}

// SMSNotifier sends notifications by SMS.
type SMSNotifier struct {
	PhoneNumber string
}

func (s SMSNotifier) Send(message string) error {
	fmt.Printf("[SMS] Envoi à %s : %s\n", s.PhoneNumber, message)
	return nil
}

// ConsoleNotifier prints notifications to the console.
type ConsoleNotifier struct{}

func (ConsoleNotifier) Send(message string) error {
	fmt.Printf("[CONSOLE] Message : %s\n", message)
	return nil
}

// Compile-time checks that each type satisfies Notifier.
var (
	_ Notifier = EmailNotifier{}
	_ Notifier = SMSNotifier{}
	_ Notifier = ConsoleNotifier{}
)
