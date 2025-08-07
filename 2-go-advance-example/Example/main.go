// main.go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Task represents a unit of work to be done.
type Task struct {
	ID      int
	Payload string
}

// Result holds the outcome of a processed task.
type Result struct {
	TaskID int
	Output string
}

// BATCH_SIZE is the number of results to collect before signaling the reporter.
const BATCH_SIZE = 5

// main is the entry point and orchestrator of our entire pipeline.
func main() {
	fmt.Println("--- Real-World Concurrent Data Processing Pipeline ---")
	rand.Seed(time.Now().UnixNano())

	// =========================================================================
	// 1. CHANNELS & ADVANCED CHANNELS: The Pipeline's Conveyor Belt
	// =========================================================================
	// WHY: Channels are the backbone of our pipeline. They allow us to safely
	// pass tasks to workers and results back from them without sharing memory
	// directly, preventing a whole class of bugs. This is the "fan-out, fan-in" pattern.
	const numTasks = 20
	tasks := make(chan Task, numTasks)     // Buffered channel for tasks (our "todo" list)
	results := make(chan Result, numTasks) // Buffered channel for results (our "done" list)

	// =========================================================================
	// 2. ATOMICS: High-Performance, Lock-Free Metrics
	// =========================================================================
	// WHY: We need to count processed tasks. We could use a mutex, but for a simple
	// counter that's updated very frequently, an atomic operation is much faster.
	// It's a single, indivisible CPU instruction, avoiding the overhead of locking.
	var processedTasks uint64

	// =========================================================================
	// 3. MUTEX: Protecting Complex Shared State
	// =========================================================================
	// WHY: Unlike a simple counter, our result map is a complex data structure.
	// Appending to a slice within a map is not an atomic operation. A Mutex
	// ensures that only one goroutine can modify this map at a time, preventing corruption.
	var resultsMutex sync.Mutex
	finalResults := make(map[int]string)

	// =========================================================================
	// 4. WAITGROUP as a BARRIER: Synchronizing Worker Start
	// =========================================================================
	// WHY: We want all our workers to be ready before we start sending them tasks.
	// A WaitGroup can act as a simple barrier. Workers will wait on it until we
	// signal them all to start at once.
	var startBarrier sync.WaitGroup
	startBarrier.Add(1) // All workers will wait until this one "Done" is called.

	// =========================================================================
	// 5. WAITGROUP for Shutdown: Ensuring All Work is Done
	// =========================================================================
	// WHY: The main goroutine needs to know when all workers have finished their
	// jobs before it can safely shut down. A WaitGroup is the perfect tool for this.
	// We increment the counter for each worker we start and each worker calls Done() on exit.
	var workerWaitGroup sync.WaitGroup

	// =========================================================================
	// 6. GOROUTINES: The Concurrent Workers
	// =========================================================================
	// WHY: This is the core of our concurrency. We spawn a pool of workers, each
	// in its own goroutine, to process tasks in parallel. This allows us to
	// utilize multiple CPU cores and handle many tasks much faster than a sequential approach.
	const numWorkers = 4
	for i := 1; i <= numWorkers; i++ {
		workerWaitGroup.Add(1)
		go worker(i, &workerWaitGroup, &startBarrier, tasks, results)
	}

	// =========================================================================
	// 7. CONDITION VARIABLE: Batching and Signaling
	// =========================================================================
	// WHY: We don't want to report every single result individually. It's more
	// efficient to batch them. A Condition Variable is ideal here. An aggregator
	// goroutine collects results, and when a specific condition is met (batch is full),
	// it uses `cond.Signal()` to wake up a reporter goroutine that was waiting for that exact moment.
	var batchMutex sync.Mutex
	batchReadyCondition := sync.NewCond(&batchMutex)
	go reporter(batchReadyCondition, &finalResults, &resultsMutex)
	go aggregator(results, &finalResults, &resultsMutex, batchReadyCondition, &processedTasks)

	// --- Pipeline Start ---
	fmt.Printf("Starting %d workers. Generating %d tasks.\n", numWorkers, numTasks)

	// Generate tasks and send them to the tasks channel.
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i, Payload: fmt.Sprintf("data for task %d", i)}
	}
	// Closing the channel is crucial! It signals to the workers (in their for...range loop)
	// that no more tasks are coming.
	close(tasks)

	// Unleash the workers!
	fmt.Println("--- All tasks sent. Releasing worker barrier. ---")
	startBarrier.Done()

	// Wait for all workers to finish processing.
	workerWaitGroup.Wait()
	fmt.Println("--- All workers have finished. ---")

	// The aggregator needs to know that no more results are coming.
	// Since the workers are done, we can safely close the results channel.
	close(results)

	// Final signal to the reporter to process any remaining items in the last batch.
	batchReadyCondition.Broadcast() // Use Broadcast to be sure.

	// A small delay to allow the final report to print.
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("\nPipeline finished. Total tasks processed: %d\n", atomic.LoadUint64(&processedTasks))
	fmt.Printf("Final collected results count: %d\n", len(finalResults))
}

// worker represents a concurrent processor in our pipeline.
func worker(id int, wg *sync.WaitGroup, startBarrier *sync.WaitGroup, tasks <-chan Task, results chan<- Result) {
	defer wg.Done()
	fmt.Printf("[Worker %d] Ready and waiting for start signal.\n", id)

	// Wait at the barrier until the main goroutine signals to start.
	startBarrier.Wait()
	fmt.Printf("[Worker %d] Starting work.\n", id)

	// The for...range on a channel is a powerful pattern. The loop will
	// automatically block until a task is available and will exit gracefully
	// when the 'tasks' channel is closed.
	for task := range tasks {
		fmt.Printf("[Worker %d] Processing Task %d...\n", id, task.ID)
		// Simulate work with a random delay
		time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)
		results <- Result{TaskID: task.ID, Output: fmt.Sprintf("Processed %s", task.Payload)}
	}

	fmt.Printf("[Worker %d] Finished. No more tasks.\n", id)
}

// aggregator collects results and signals the reporter when a batch is ready.
func aggregator(results <-chan Result, finalResults *map[int]string, resultsMutex *sync.Mutex, batchReady *sync.Cond, processedCounter *uint64) {
	batchCount := 0
	for result := range results {
		// Use the mutex to safely write to the shared map.
		resultsMutex.Lock()
		(*finalResults)[result.TaskID] = result.Output
		resultsMutex.Unlock()

		// Use atomic for the simple counter.
		atomic.AddUint64(processedCounter, 1)

		batchCount++
		// =========================================================================
		// 8. SELECT (inside aggregator, though often in worker)
		// =========================================================================
		// WHY: While not a complex select here, this logic block demonstrates a choice.
		// If the batch is full, we signal. Otherwise, we continue. A more complex
		// worker might use `select` to choose between a new task or a shutdown signal.
		if batchCount >= BATCH_SIZE {
			fmt.Println("--- AGGREGATOR: Batch is full. Signaling reporter. ---")
			batchReady.Signal() // Wake up the waiting reporter.
			batchCount = 0      // Reset for the next batch.
		}
	}
	fmt.Println("--- AGGREGATOR: Result channel closed. Shutting down. ---")
}

// reporter waits for a signal that a batch is ready and then "reports" it.
func reporter(batchReady *sync.Cond, finalResults *map[int]string, resultsMutex *sync.Mutex) {
	fmt.Println("[REPORTER] Waiting for a full batch...")
	for {
		batchReady.L.Lock()
		// This is the key to condition variables: we wait in a loop.
		// The Wait() call atomically unlocks the mutex and puts the goroutine to sleep.
		// When woken up, it re-acquires the lock before continuing.
		batchReady.Wait()

		resultsMutex.Lock()
		fmt.Printf("\n>>> [REPORTER] Woke up! Reporting on %d results. <<<\n", len(*finalResults))
		// In a real system, this would send the batch to another service,
		// write to a database, etc. Here we just print a summary.
		for id, res := range *finalResults {
			fmt.Printf("  - Task %d: %s\n", id, res)
		}
		fmt.Println(">>> [REPORTER] Report finished. Going back to sleep. <<<")
		resultsMutex.Unlock()

		batchReady.L.Unlock()
	}
}
