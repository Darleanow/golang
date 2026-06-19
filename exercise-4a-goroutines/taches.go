package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// simulerTravail sleeps for a random duration up to maxMs milliseconds.
// Keeping maxMs low in exercice1 makes the premature exit easier to observe.
func simulerTravail(maxMs int) {
	time.Sleep(time.Duration(rand.Intn(maxMs)) * time.Millisecond)
}

// tacheNonSync runs a task without any synchronisation (Exercice 1).
func tacheNonSync(id int) {
	fmt.Printf("Goroutine %d : début de la tâche...\n", id)
	simulerTravail(300)
	fmt.Printf("Goroutine %d : tâche terminée.\n", id)
}

// tacheAvecWaitGroup signals its completion through the WaitGroup with a
// deferred Done, so it is counted whether it returns normally or panics
// (Exercice 2).
func tacheAvecWaitGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Goroutine %d : début de la tâche...\n", id)
	simulerTravail(400)
	fmt.Printf("Goroutine %d : tâche terminée.\n", id)
}

// tacheAvecCanal sends its result on resultats once the work is done, then
// signals completion (Exercice 3).
func tacheAvecCanal(id int, resultats chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Goroutine %d : début de la tâche...\n", id)
	simulerTravail(600)
	resultats <- fmt.Sprintf("Goroutine %d a terminé avec succès.", id)
}

// travailleur consumes task IDs from taches until the channel is closed,
// emitting one result per task. The deferred Done releases the WaitGroup when
// the input channel drains (Exercice 4).
func travailleur(id int, taches <-chan int, resultats chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for tache := range taches {
		simulerTravail(500)
		resultats <- fmt.Sprintf("Travailleur %d a traité la tâche %d.", id, tache)
	}
}
