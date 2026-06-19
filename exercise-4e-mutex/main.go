// Command exercise-4e-mutex demonstrates the race condition that arises from
// unprotected concurrent increments and shows how sync.Mutex and sync/atomic
// fix it.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	nbGoroutines           = 100
	incrementsParGoroutine = 1000
	valeurAttendue         = nbGoroutines * incrementsParGoroutine
)

var compteurNonSync int

func incrementerNonSync(wg *sync.WaitGroup) {
	defer wg.Done()
	for range incrementsParGoroutine {
		compteurNonSync++
	}
}

func etape1() {
	fmt.Println("Étape 1 - Sans synchronisation (race condition)")
	compteurNonSync = 0
	var wg sync.WaitGroup
	for range nbGoroutines {
		wg.Add(1)
		go incrementerNonSync(&wg)
	}
	wg.Wait()
	fmt.Printf("  Résultat : %d  (attendu : %d)  - correct : %v\n\n",
		compteurNonSync, valeurAttendue, compteurNonSync == valeurAttendue)
}

var (
	compteurMutex int
	mu            sync.Mutex
)

func incrementerMutex(wg *sync.WaitGroup) {
	defer wg.Done()
	for range incrementsParGoroutine {
		mu.Lock()
		compteurMutex++
		mu.Unlock()
	}
}

func etape2() {
	fmt.Println("Étape 2 - Avec sync.Mutex")
	compteurMutex = 0
	var wg sync.WaitGroup
	for range nbGoroutines {
		wg.Add(1)
		go incrementerMutex(&wg)
	}
	wg.Wait()
	fmt.Printf("  Résultat : %d  (attendu : %d)  - correct : %v\n\n",
		compteurMutex, valeurAttendue, compteurMutex == valeurAttendue)
}

// atomic.AddInt64 is faster than a mutex for a single-variable critical
// section but cannot protect multi-step logic - the right tool depends on the
// complexity of the guarded block.
var compteurAtomic int64

func incrementerAtomic(wg *sync.WaitGroup) {
	defer wg.Done()
	for range incrementsParGoroutine {
		atomic.AddInt64(&compteurAtomic, 1)
	}
}

func etape3() {
	fmt.Println("Étape 3 - Avec sync/atomic (bonus)")
	atomic.StoreInt64(&compteurAtomic, 0)
	var wg sync.WaitGroup
	for range nbGoroutines {
		wg.Add(1)
		go incrementerAtomic(&wg)
	}
	wg.Wait()
	v := atomic.LoadInt64(&compteurAtomic)
	fmt.Printf("  Résultat : %d  (attendu : %d)  - correct : %v\n",
		v, valeurAttendue, v == int64(valeurAttendue))
}

func main() {
	etape1()
	etape2()
	etape3()
}
