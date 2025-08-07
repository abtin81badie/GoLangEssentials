package concurrency

import (
	"fmt"
	"time"
)

// worker is a function that will run concurrently.
// It receives a channel to send its result back on.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // Simulate work
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

// DemonstrateGoroutinesAndChannels shows a basic worker pool pattern.
func DemonstrateGoroutinesAndChannels() {
	fmt.Println("\n[Goroutines and Channels]")

	numJobs := 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start 3 concurrent workers. They are blocked because the `jobs` channel is empty.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs to the `jobs` channel. The workers will pick them up.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close the channel to signal no more jobs are coming.

	// Wait for all the results.
	for a := 1; a <= numJobs; a++ {
		<-results // Block until a result is received.
	}
	fmt.Println("All jobs are done.")
}
