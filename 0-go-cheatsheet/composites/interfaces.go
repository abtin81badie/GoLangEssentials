package composites

import "fmt"

// Speaker is an interface. It defines a set of methods.
// Any type that implements all methods for an interface implicitly satisfies that interface.
type Speaker interface {
	Speak() string
}

// Dog is a struct that will implement the Speaker interface.
type Dog struct {
	Name string
}

// Speak is a method that satisfies the Speaker interface for Dog.
func (d Dog) Speak() string {
	return "Woof!"
}

// Human is another struct that implements the Speaker interface.
type Human struct {
	Name string
}

// Human's Speak method satisfies the Speaker interface.
func (h Human) Speak() string {
	return "Hello!"
}

// MakeItSpeak is a function that takes a Speaker interface.
func MakeItSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

// DemonstrateInterfaces shows how to define and use interfaces in Go.
func DemonstrateInterfaces() {
	fmt.Println("\n[Interfaces]")

	dog := Dog{Name: "Buddy"}
	human := Human{Name: "Alice"}

	// Both Dog and Human implement the Speaker interface.
	fmt.Print("A dog says: ")
	MakeItSpeak(dog) // Output: A dog says: Woof!

	fmt.Print("A human says: ")
	MakeItSpeak(human) // Output: A human says: Hello!

	// The empty interface `interface{}` (or `any`) can hold any values.
	// Use type assertion to get underlying type.
	var i any = "a string value"
	s, ok := i.(string) // Type assertion to string
	if ok {
		fmt.Printf("Empty interface hold a string: '%s\n", s)
	}
}
