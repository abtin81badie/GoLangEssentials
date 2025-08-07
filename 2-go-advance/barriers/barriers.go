// barriers/barriers.go
package barriers

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// THEORY:
// A "barrier" is a synchronization point that a group of goroutines must all
// reach before any are allowed to proceed. A WaitGroup can be used for a single-use barrier.

// RunExample demonstrates using two WaitGroups to create two synchronization points.
func RunExample() {
	fmt.Println("\n--- Barrier Example ---")

	const numWorkers = 5
	var phase1, phase2 sync.WaitGroup

	phase1.Add(numWorkers)
	phase2.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go func(id int) {
			fmt.Printf("[Worker %d] Phase 1...\n", id)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			fmt.Printf("[Worker %d] Finished Phase 1.\n", id)
			phase1.Done()

			phase1.Wait() // Barrier

			fmt.Printf("[Worker %d] Phase 2...\n", id)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			fmt.Printf("[Worker %d] Finished Phase 2.\n", id)
			phase2.Done()
		}(i)
	}

	fmt.Println("Main: Waiting for all phases.")
	phase2.Wait()
	fmt.Println("Main: All phases complete.")
}
