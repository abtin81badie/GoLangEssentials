package selects

import (
	"fmt"
	"time"
)

// THEORY:
// The `select` statement lets a goroutine wait on multiple channel operations.
// It blocks until one of its cases can run. If multiple are ready, it chooses one randomly.
// A `default` case makes it non-blocking.

// RunExample demonstrates waiting on multiple channels and implementing a timeout.
func RunExample() {
	fmt.Println("\n--- Select Example ---")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("Waiting for a message...")
		select {
		case msg1 := <-c1:
			fmt.Println("Received:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received:", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout: No message received.")
		}
	}
}
