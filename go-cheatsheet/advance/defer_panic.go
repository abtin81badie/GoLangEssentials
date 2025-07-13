package advance

import "fmt"

func cleanup() {
	// A deferred function runs after the surrounding function returns.
	// It's often used for cleanup, like closing files.
	fmt.Println("Performing cleanup.")
}

// DemonstrateDefer shows how defer works.
func DemonstrateDefer() {
	fmt.Println("\n[Defer]")

	defer cleanup() // This will run last.
	fmt.Println("Doing some work...")
	// The `panic` keyword causes the program to stop immediately, but deferred functions still run.
	// We comment it out here so the program can continue.
	// panic("something terrible happened")
}
