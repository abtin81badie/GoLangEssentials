package raceconditions

import (
	"fmt"
	"sync"
)

// THEORY:
// A "race condition" occurs when multiple goroutines access the same shared data concurrently,
// and at least one of them modifies it. The result is unpredictable.
// Detect race conditions with the -race flag: `go run -race main.go`

// RunExample demonstrates a race condition.
func RunExample() {
	fmt.Println("\n--- Race Condition Example ---")
	fmt.Println("Run this with 'go run -race main.go' to see the race detector in action.")

	var wg sync.WaitGroup
	var counter int = 0

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter++
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Expected counter value: 10000, Got: %d\n", counter)
}
