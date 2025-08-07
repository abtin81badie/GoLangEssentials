package goroutines

import (
	"fmt"
	"time"
)

// THEORY:
// A goroutine is a lightweight thread of execution managed by the Go runtime.
// Starting a new goroutine is as simple as using the 'go' keyword before a function call.
// Goroutines run in the same address space, so access to shared memory must be synchronized.
// When you start a goroutine, the main program flow does not wait for it to finish.

func sayHello() {
	fmt.Println("Hello from the goroutine!")
}

// RunExample demonstrates the most basic form of concurrency in Go.
func RunExample() {
	fmt.Println("\n--- Goroutine Example ---")

	go sayHello()

	fmt.Println("Hello from the main function!")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main function finished.")
}
