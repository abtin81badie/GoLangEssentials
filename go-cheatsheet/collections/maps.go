package collections

import "fmt"

// DemonstrateMaps shows how to use maps in Go, which are key-value pairs. (hash tables)
func DemonstrateMaps() {
	fmt.Println("\n[Maps]")

	// Maps are key-value pairs, similar to dictionaries in Python or hash tables in other languages.
	// create with make or map literal
	ages := make(map[string]int) // Create an empty map with string keys and int values
	ages["Alice"] = 30         // Add a key-value pair
	ages["Bob"] = 25           // Add another key-value pair
	fmt.Println("Ages map:", ages) // Output: Ages map: map[Alice:30 Bob:25]

	// Getting a value by key
	fmt.Println("Alice's age:", ages["Alice"]) // Output: Alice's age: 30

	// Deleting a key-value pair
	delete(ages, "Bob") // Remove Bob from the map
	fmt.Println("After deleting Bob, Ages map:", ages) // Output: After deleting Bob, Ages map: map[Alice:30]

	// Checking if a key exists
	if age, exists := ages["Bob"]; exists {
		fmt.Println("Bob's age:", age) // This won't execute since Bob was deleted
	} else {
		fmt.Println("Bob does not exist in the map") // Output: Bob does not exist in the map
	}

	// 'range' iterates over key-value pairs in a map.
	fmt.Println("Iterating over Ages map:")
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age) // Output: Alice is 30 years old
	}
}