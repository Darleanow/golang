// Command exercise-4d-select demonstrates select-based event multiplexing
// through a simulated monitoring system.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// dataProducer and alertProducer both accept a quit channel so they exit
// cleanly when the shutdown signal fires - no goroutine leak after main returns.
func dataProducer(dataChannel chan<- string, quit <-chan struct{}) {
	for {
		delai := time.Duration(1+rand.Intn(3)) * time.Second
		select {
		case <-quit:
			return
		case <-time.After(delai):
			dataChannel <- fmt.Sprintf("Température : %d°C", 20+rand.Intn(15))
		}
	}
}

func alertProducer(alertChannel chan<- string, quit <-chan struct{}) {
	for {
		delai := time.Duration(5+rand.Intn(6)) * time.Second
		select {
		case <-quit:
			return
		case <-time.After(delai):
			alertChannel <- "Niveau critique atteint !"
		}
	}
}

func main() {
	dataChannel := make(chan string)
	alertChannel := make(chan string)
	quitChannel := make(chan struct{})

	// NewTicker must be stopped explicitly; defer guarantees it even if main
	// returns early through the quit case.
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	go dataProducer(dataChannel, quitChannel)
	go alertProducer(alertChannel, quitChannel)

	go func() {
		time.Sleep(15 * time.Second)
		close(quitChannel)
	}()

	fmt.Println("Système de surveillance démarré.")

	for {
		select {
		case msg := <-dataChannel:
			fmt.Printf("[MESURE] %s\n", msg)
		case msg := <-alertChannel:
			fmt.Printf("[ALERTE CRITIQUE] %s\n", msg)
		case <-ticker.C:
			fmt.Println("[STATUS] Vérification système...")
		case <-quitChannel:
			fmt.Println("Signal d'arrêt reçu. Arrêt du système.")
			return
		}
	}
}
