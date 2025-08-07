// advancedchannels/advancedchannels.go
package advancedchannels

import (
	"fmt"
	"sync"
	"time"
)

// THEORY:
// Advanced channel patterns allow for building complex concurrent data pipelines.
// - Buffered Channels: `make(chan T, capacity)` creates a channel with a buffer.
//   This decouples the sender and receiver, as sends only block when the buffer is full.
// - Closing Channels: `close(ch)` is used by a sender to indicate that no more
//   values will be sent on a channel. This is crucial for receivers to know when to stop reading.
// - Fan-out, Fan-in: A common and powerful pattern where a single producer goroutine
//   distributes (fans-out) work to multiple worker goroutines. These workers then send their
//   results to a single channel, which another goroutine collects (fans-in).

// RunExample demonstrates a fan-out, fan-in pipeline to process a number of "jobs".
func RunExample() {
	fmt.Println("\n--- Advanced Channel Patterns (Fan-out, Fan-in) ---")

	const numJobs = 10
	const numWorkers = 3

	// The 'jobs' channel is used to send work from the producer to the workers.
	// It is buffered to hold all jobs at once, though this isn't strictly necessary.
	jobs := make(chan int, numJobs)

	// The 'results' channel is used by workers to send their results back.
	results := make(chan int, numJobs)

	// A WaitGroup is used to wait for all worker goroutines to finish.
	var wg sync.WaitGroup

	// --- FAN-OUT ---
	// This loop launches our pool of worker goroutines.
	// All of them read from the SAME `jobs` channel, distributing the work automatically.
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		// Each worker is a goroutine.
		go func(id int, jobs <-chan int, results chan<- int) {
			// Decrement the WaitGroup counter when the goroutine finishes.
			defer wg.Done()

			// This `for...range` loop will automatically receive values from the 'jobs'
			// channel until it is closed and empty. This is the idiomatic way to process
			// work from a channel.
			for j := range jobs {
				fmt.Printf("Worker %d started job %d\n", id, j)
				time.Sleep(time.Second) // Simulate doing some work
				fmt.Printf("Worker %d finished job %d\n", id, j)
				// Send the result of the work to the 'results' channel.
				results <- j * 2
			}
		}(w, jobs, results)
	}

	// --- PRODUCER ---
	// This section sends all the jobs to the `jobs` channel.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// It's crucial to close the `jobs` channel after sending all the work.
	// This signals to the worker goroutines (in their `for...range` loops) that
	// there's no more work, allowing them to terminate gracefully.
	close(jobs)

	// --- FAN-IN (Part 1: The Waiter) ---
	// We need a way to know when all workers have finished so we can safely close
	// the 'results' channel.
	// We launch a separate goroutine that waits for all workers to complete (using the WaitGroup)
	// and *then* closes the 'results' channel. This prevents a premature exit from our final collection loop.
	go func() {
		wg.Wait()
		close(results)
	}()

	// --- FAN-IN (Part 2: The Collector) ---
	// This final loop collects all the results from the workers.
	// It will block until a result is available, and the loop will terminate
	// automatically when the 'results' channel is closed by our "waiter" goroutine above.
	total := 0
	for result := range results {
		total += result
	}

	fmt.Printf("All jobs are done. Total of results: %d\n", total)
}
