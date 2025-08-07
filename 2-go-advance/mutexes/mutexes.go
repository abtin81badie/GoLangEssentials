package mutexes

import (
	"fmt"
	"sync"
)

// THEORY:
// A Mutex (Mutual Exclusion lock) prevents race conditions by ensuring only one
// goroutine can access a "critical section" of code at a time.
// - Lock(): Acquires the lock, blocking if it's held by another goroutine.
// - Unlock(): Releases the lock.

// RunExample fixes the race condition using a Mutex.
func RunExample() {
	fmt.Println("\n--- Mutex Example ---")

	var wg sync.WaitGroup
	var mutex sync.Mutex
	var counter int = 0

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Expected counter value: 10000, Got: %d\n", counter)
}
