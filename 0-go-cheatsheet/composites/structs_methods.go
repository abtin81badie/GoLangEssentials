package composites

import "fmt"

// Person is a struct, a collection of named fields.
type Person struct {
	Name string
	Age  int
}

// Greet is a "method" associated with the Person struct.
func (p Person) Greet() string {
	return "Hello, my name is " + p.Name + " and I am " + fmt.Sprint(p.Age) + " years old."
}

// HaveBirthday is a method with a pointer receiver (*Person), allowing it to modify the original struct.
func (p *Person) HaveBirthday() {
	p.Age++ // Increment the Age field by 1
}

// DemonstrateStructsMethods shows how to define and use structs and methods in Go.
func DemonstrateStructsMethods() {
	fmt.Println("\n[Structs & Methods]")

	// Create a new Person instance
	alice := Person{Name: "Alice", Age: 30}

	// Call the Greet method
	fmt.Println(alice.Greet()) // Output: Hello, my name is Alice and I am 30 years old.

	// Call the HaveBirthday method to increment Alice's age
	alice.HaveBirthday()
	fmt.Println("After birthday:", alice.Greet()) // Output: After birthday: Hello, my name is Alice and I am 31 years old.
}
