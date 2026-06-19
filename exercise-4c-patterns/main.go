// Command exercise-4c-patterns demonstrates the Worker Pool and Fan-out/Fan-in
// concurrency patterns.
package main

import (
	"fmt"
	"sync"
	"time"
)

type resultat struct {
	nombre int
	somme  int
}

func sumDivisors(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

// generateNumbers closes jobs after sending all values so worker range loops
// terminate naturally without a separate stop signal.
func generateNumbers(numJobs int, jobs chan<- int) {
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)
}

func worker(id int, jobs <-chan int, results chan<- resultat, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs {
		r := resultat{nombre: n, somme: sumDivisors(n)}
		fmt.Printf("  Worker %d : sumDivisors(%d) = %d\n", id, n, r.somme)
		results <- r
	}
}

func main() {
	const numWorkers = 4
	const numJobs = 20

	debut := time.Now()

	jobs := make(chan int, numJobs)
	results := make(chan resultat, numJobs)

	var wg sync.WaitGroup

	for id := 1; id <= numWorkers; id++ {
		wg.Add(1)
		go worker(id, jobs, results, &wg)
	}

	go generateNumbers(numJobs, jobs)

	// Closing results from a goroutine lets main drain it with range while
	// workers are still running — avoids a separate collection phase after Wait.
	go func() {
		wg.Wait()
		close(results)
	}()

	total := 0
	for r := range results {
		total += r.somme
	}

	fmt.Printf("\nSomme totale des diviseurs (1 à %d) : %d\n", numJobs, total)
	fmt.Printf("Temps d'exécution avec %d workers : %s\n", numWorkers, time.Since(debut))
}
