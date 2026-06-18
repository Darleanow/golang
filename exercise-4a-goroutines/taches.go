package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// simulerTravail blocks for a random duration between 50 and 500 ms to mimic
// some real work. The global generator is auto-seeded since Go 1.20, so there
// is no need for the deprecated rand.Seed.
func simulerTravail() {
	time.Sleep(time.Duration(50+rand.Intn(451)) * time.Millisecond)
}

// tacheNonSync runs a task without any synchronisation (Exercice 1).
func tacheNonSync(id int) {
	fmt.Printf("Goroutine %d : début de la tâche...\n", id)
	simulerTravail()
	fmt.Printf("Goroutine %d : tâche terminée.\n", id)
}

// tacheAvecWaitGroup signals its completion through the WaitGroup with a
// deferred Done, so it is counted whether it returns normally or panics
// (Exercice 2).
func tacheAvecWaitGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Goroutine %d : début de la tâche...\n", id)
	simulerTravail()
	fmt.Printf("Goroutine %d : tâche terminée.\n", id)
}

// tacheAvecCanal sends its result on resultats once the work is done, then
// signals completion (Exercice 3).
func tacheAvecCanal(id int, resultats chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Goroutine %d : début de la tâche...\n", id)
	simulerTravail()
	resultats <- fmt.Sprintf("Goroutine %d a terminé avec succès.", id)
}

// travailleur consumes task IDs from taches until the channel is closed,
// emitting one result per task. The deferred Done releases the WaitGroup when
// the input channel drains (Exercice 4).
func travailleur(id int, taches <-chan int, resultats chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for tache := range taches {
		simulerTravail()
		resultats <- fmt.Sprintf("Travailleur %d a traité la tâche %d.", id, tache)
	}
}
