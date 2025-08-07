package advance

import "fmt"

// DemonstratePointers explains how pointers store memory addresses.
func DemonstratePointers() {
	fmt.Println("\n[Pointers]")

	val := 42

	// `p` is a pointer to `val`. The `&` operator gets the address of `val`.
	p := &val

	fmt.Println("Original value: %d\n", val)
	fmt.Println("Pointer address: %v\n", p)

	// The `*` operator "dereference" the pointer, giving the value at the address.
	fmt.Printf("Value via pointer: %d\n", *p)

	// We can change the original value through the pointer.
	*p = 100
	fmt.Printf("Value after changing via pointer: %d\n", val)
}
