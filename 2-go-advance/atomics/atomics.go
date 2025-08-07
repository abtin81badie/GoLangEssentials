package atomics

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// THEORY:
// The `sync/atomic` package provides low-level atomic memory primitives that are
// managed directly by the hardware. These operations, such as adding to an integer
// or swapping a value, are guaranteed to be "atomic," meaning they are performed
// in a single, indivisible step without interruption from other goroutines.
//
// For simple operations like incrementing a counter or updating a flag, atomic
// functions can be significantly faster than using a mutex. A mutex involves a
// more complex mechanism of locking and unlocking that can cause goroutines to
// block and be rescheduled by the Go runtime. Atomic operations, by contrast,
// are typically single CPU instructions and do not involve blocking.

// RunExample demonstrates safe concurrent counting using atomic operations.
func RunExample() {
	fmt.Println("\n--- Atomic Operations Example ---")

	var wg sync.WaitGroup
	// Our counter must be an unsigned 64-bit integer to use the atomic functions.
	var counter uint64 = 0

	// We'll launch 100 goroutines that all increment the same counter.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Each goroutine will increment the counter 100 times.
			for j := 0; j < 100; j++ {
				// Atomically add 1 to the counter. This is safe to do from
				// multiple goroutines concurrently without a race condition.
				// The first argument must be a pointer to the variable.
				atomic.AddUint64(&counter, 1)
			}
		}()
	}

	wg.Wait()

	// To safely read the value of the counter, we should also use an atomic operation.
	// `atomic.LoadUint64` safely retrieves the value.
	finalCount := atomic.LoadUint64(&counter)

	// With atomic operations, the final value will always be the expected 10000.
	fmt.Printf("Expected counter value: 10000, Got: %d\n", finalCount)
}
