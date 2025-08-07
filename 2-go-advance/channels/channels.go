package channels

import (
	"fmt"
	"time"
)

// THEORY:
// A channel is a typed conduit for sending and receiving values. It's the idiomatic
// way to communicate and synchronize between goroutines in Go.
// `ch <- v` sends, `v := <-ch` receives. Sends/receives block by default.

// RunExample demonstrates basic channel communication.
func RunExample() {
	fmt.Println("\n--- Channel Example ---")

	messages := make(chan string)

	go func() {
		fmt.Println("Goroutine: preparing to send...")
		time.Sleep(1 * time.Second)
		messages <- "Hello from the channel!"
		fmt.Println("Goroutine: message sent.")
	}()

	fmt.Println("Main: waiting to receive...")
	msg := <-messages
	fmt.Printf("Main: received message: '%s'\n", msg)
}
