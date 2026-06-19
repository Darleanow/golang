// Command exercise-4a-goroutines launches and synchronises goroutines with
// WaitGroup, channels and a worker pool.
package main

import (
	"fmt"
	"sync"
	"time"
)

// exercice1 launches goroutines without synchronisation. The trailing Sleep is
// only a crude way to keep the program alive long enough to see their output:
// without it (or a WaitGroup) main would return and the goroutines could be
// killed before finishing. The proper fix is Exercice 2.
func exercice1() {
	fmt.Println("Exercice 1 - Lancement sans synchronisation")

	for id := 1; id <= 5; id++ {
		go tacheNonSync(id)
	}
	fmt.Println("Toutes les goroutines lancées.")

	// 600 ms > 300 ms max from simulerTravail, so goroutines will likely finish
	// but it is not guaranteed. WaitGroup in exercice2 is the real fix.
	time.Sleep(600 * time.Millisecond)
}

// exercice2 waits for every goroutine through a WaitGroup before returning.
func exercice2() {
	fmt.Println("\nExercice 2 - Synchronisation avec sync.WaitGroup")

	var wg sync.WaitGroup
	for id := 1; id <= 5; id++ {
		wg.Add(1)
		go tacheAvecWaitGroup(id, &wg)
	}
	fmt.Println("Toutes les goroutines lancées.")

	wg.Wait()
	fmt.Println("Toutes les goroutines ont terminé leur exécution.")
}

// exercice3 collects the goroutines' results through a channel. The channel is
// buffered so the workers can send without blocking before main drains it after
// wg.Wait(); an unbuffered channel would deadlock with this read-after-wait order.
func exercice3() {
	fmt.Println("\nExercice 3 - Communication par canaux")

	const nbTaches = 5
	resultats := make(chan string, nbTaches)

	var wg sync.WaitGroup
	for id := 1; id <= nbTaches; id++ {
		wg.Add(1)
		go tacheAvecCanal(id, resultats, &wg)
	}

	wg.Wait()
	close(resultats)

	for msg := range resultats {
		fmt.Println(msg)
	}
}

// exercice4 dispatches a batch of tasks to a fixed pool of workers. Closing the
// taches channel ends the workers' range loops; closing resultats after
// wg.Wait() lets main drain the collected results.
func exercice4() {
	fmt.Println("\nExercice 4 - Pool de travailleurs")

	const nbTravailleurs = 3
	const nbTaches = 10

	taches := make(chan int, nbTaches)
	resultats := make(chan string, nbTaches)

	var wg sync.WaitGroup
	for id := 1; id <= nbTravailleurs; id++ {
		wg.Add(1)
		go travailleur(id, taches, resultats, &wg)
	}

	for t := 1; t <= nbTaches; t++ {
		taches <- t
	}
	close(taches)

	wg.Wait()
	close(resultats)

	for msg := range resultats {
		fmt.Println(msg)
	}
}

func main() {
	exercice1()
	exercice2()
	exercice3()
	exercice4()
}
