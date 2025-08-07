package condvars

import (
	"fmt"
	"sync"
	"time"
)

// THEORY:
// A Condition Variable (sync.Cond) allows goroutines to wait for a specific condition to become true.
// It's always associated with a Mutex.
// - Wait(): Atomically unlocks the mutex and suspends the goroutine.
// - Signal(): Wakes up ONE waiting goroutine.
// - Broadcast(): Wakes up ALL waiting goroutines.

// RunExample demonstrates a producer-consumer scenario using sync.Cond.
func RunExample() {
	fmt.Println("\n--- Condition Variable Example ---")

	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)
	var queue []int

	go func() { // Consumer
		fmt.Println("Consumer: Ready.")
		cond.L.Lock()
		for len(queue) == 0 {
			fmt.Println("Consumer: Waiting...")
			cond.Wait()
			fmt.Println("Consumer: Woke up.")
		}
		item := queue[0]
		queue = queue[1:]
		fmt.Printf("Consumer: Consumed item %d\n", item)
		cond.L.Unlock()
	}()

	go func() { // Producer
		fmt.Println("Producer: Working...")
		time.Sleep(2 * time.Second)
		cond.L.Lock()
		item := 42
		queue = append(queue, item)
		fmt.Printf("Producer: Produced item %d, signaling.\n", item)
		cond.Signal()
		cond.L.Unlock()
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Condition variable example finished.")
}
