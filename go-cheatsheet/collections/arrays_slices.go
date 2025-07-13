package collections

import "fmt"

// DemonstrateArraysAndSlices shows fixed-size arrays and dynamic slices in Go.
func DemonstrateArraysAndSlices() {
	fmt.Println("\n[Arrays & Slices]")

	// ARRAYS have a fixed size defined at compile time.
	var a [3]int                 // Array of 3 integers, initialized to zero values
	a[1] = 10                    // Set the second element to 10
	fmt.Println("Array 'a':", a) // Output: Array a: [0 10 0]

	// SLICES are dynamic, more common than arrays and can grow or shrink.
	// Create a slice with initial values
	s := make([]string, 3) // Slice of strings with length 3
	s[0] = "GO"
	s[1] = "is"
	s[2] = "fun"
	fmt.Println("Slice 's':", s) // Output: Slice s: [GO is fun]

	// Append adds elements and may resize the underlying array.
	s = append(s, "and", "powerful")          // Append two more elements
	fmt.Println("Slice 's' after append:", s) // Output: Slice s after append: [GO is fun and powerful]

	// 'range' iterates over elements in a slice. (and other collections)
	fmt.Print("Range over slice 's': ")
	for i, v := range s {
		fmt.Printf("Index %d: %s ", i, v)
	}
	fmt.Println() // New line after range output
}
