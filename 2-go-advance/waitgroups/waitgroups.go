package waitgroups

import (
	"fmt"
	"sync"
	"time"
)

// THEORY:
// A sync.WaitGroup is a common way to wait for a collection of goroutines to finish executing.
// It acts as a counter.
// - Add(delta int): Increases the counter by 'delta'.
// - Done(): Decreases the counter by one.
// - Wait(): Blocks until the counter becomes zero.

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

// RunExample shows how to use a WaitGroup to manage multiple goroutines.
func RunExample() {
	fmt.Println("\n--- WaitGroup Example ---")
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	fmt.Println("Main goroutine is waiting for workers to finish...")
	wg.Wait()
	fmt.Println("All workers have finished. Main function exiting.")
}
