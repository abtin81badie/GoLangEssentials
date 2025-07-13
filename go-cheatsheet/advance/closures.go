package advance

import "fmt"

// intSequence returns another function, which we define anonymously in the body of intSequence.
// the returned function is a "closure" that captures the variable `i`.
func intSequence() func() int {
	i := 0 // This variable is captured by the closure
	return func() int {
		i++ // Increment i each time the returned function is called
		return i
	}
}

// DemonstrateClosures shows how closures work.
func DemonstrateClosures() {
	fmt.Println("\n[Closures]")

	// nextInt is a function, but it has its own state for `i`.
	nextInt := intSequence()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	// Each call to intSequence creates a new, independent closure.
	newInts := intSequence()
	fmt.Println(newInts()) // 1
}
