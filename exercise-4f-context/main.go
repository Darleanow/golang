// Command exercise-4f-context demonstrates context-based cancellation and
// timeout propagation through a simulated long-running operation.
package main

import (
	"context"
	"fmt"
	"time"
)

// effectuerOperationLongue checks ctx.Done() between every step so it can
// abort early without leaking the goroutine - a naked time.Sleep would ignore
// the cancellation signal until the sleep expires.
func effectuerOperationLongue(ctx context.Context, id string) error {
	fmt.Printf("[%s] Début de l'opération...\n", id)
	for etape := 1; etape <= 5; etape++ {
		select {
		case <-ctx.Done():
			fmt.Printf("[%s] Opération annulée à l'étape %d : %v\n", id, etape, ctx.Err())
			return ctx.Err()
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("[%s] Traitement étape %d...\n", id, etape)
		}
	}
	fmt.Printf("[%s] Opération terminée avec succès.\n", id)
	return nil
}

// runAvecTimeout wraps the call so main can compare two timeout scenarios
// without duplicating the context/channel boilerplate.
func runAvecTimeout(timeout time.Duration, label string) {
	fmt.Printf("\n--- %s (timeout : %s) ---\n", label, timeout)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Buffered so the goroutine never blocks on send even if we return early.
	resultChan := make(chan error, 1)
	go func() {
		resultChan <- effectuerOperationLongue(ctx, label)
	}()

	// The goroutine always sends on resultChan (error or nil), so that case
	// fires first. ctx.Done() here is a safety net: it triggers if the parent
	// context is cancelled externally before the goroutine finishes.
	select {
	case err := <-resultChan:
		if err != nil {
			fmt.Printf("Main [%s] : opération annulée : %v\n", label, err)
		} else {
			fmt.Printf("Main [%s] : opération terminée avec succès.\n", label)
		}
	case <-ctx.Done():
		fmt.Printf("Main [%s] : contexte annulé avant la fin de l'opération : %v\n", label, ctx.Err())
	}
}

func main() {
	fmt.Println("Démarrage du programme principal.")

	// 2 s < 5×500 ms = 2.5 s: the timeout fires before the operation finishes.
	runAvecTimeout(2*time.Second, "Tâche courte (2s)")

	// 3 s > 2.5 s: all steps complete before the deadline.
	runAvecTimeout(3*time.Second, "Tâche longue (3s)")

	fmt.Println("\nFin du programme principal.")
}
