package basics

import "fmt"

// DemonstrateVariables shows various ways to declare and initialize variables and constants.
func DemonstrateVariables() {
	fmt.Println("\n[Variables & Constants]")

	// 'var' declares a variable. It's zero-valued until assigned.
	var a int
	fmt.Printf("var a int -> a: %d (zero value)\n", a)

	// 'var' with an initializer. Type is inferred.
	var b = "hello"
	fmt.Printf("var b = \"hello\" -> b: %s\n", b)

	// Short declaration `:=` infers type and initializes. Only for new variables.
	c := true
	fmt.Printf("c := true -> c: %t\n", c)

	// Constants are declared with 'const' and cannot be changed.
	const Pi = 3.14159
	fmt.Printf("const Pi = 3.14159 -> Pi: %f\n", Pi)
}
