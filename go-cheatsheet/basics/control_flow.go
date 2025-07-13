package basics

import "fmt"

// DemonstrateControlFlow shows conditional logic and loops in Go.
func DemonstrateControlFlow() {
	fmt.Println("\n[Control Flow]")

	// If-else statement
	if num := 9; num < 0 {
		fmt.Println("Number is negative.")
	} else if num < 10 {
		fmt.Println("Number is single-digit positive.")
	} else {
		fmt.Println("Number is double-digit or larger.")
	}

	// For loop (the only loop in Go)
	for i := 0; i < 5; i++ {
		fmt.Printf("For loop iteration: %d\n", i)
	}

	// switch statement
	switch day := "Monday"; day {
		case "Monday","Tuesday", "Wednesday", "Thursday", "Friday":
			fmt.Println("It's a weekday.")
		case "Saturday", "Sunday":
			fmt.Println("It's the weekend.")
		default:
			fmt.Println("Unknown day.")
	}
}
