package basics

import "fmt"

// passValue demonstrates that the original value is not modified when passed by value.
func passValue(x int) {
	x = 10 // This change won't affect the original variable.
	fmt.Println("Inside passValue, x:", x)
}

//passByReference demonstrates that the original value is modified when passed by reference.
func passByReference(x *int) {
	*x = 10 // This change will affect the original variable.
	fmt.Println("Inside passByReference, x:", *x)
}

// multiReturn demonstrates a function that returns multiple values.
func multiReturn(isGood bool) (string, error) {
	if isGood {
		return "Everything is fine!", nil
	}
	return "", fmt.Errorf("an error occurred")
}

// DemmonstrateFunctions runs examples of different function types in Go.
func DemonstrateFunctions() {
	fmt.Println("\n[Functions]")

	// Function with value parameter
	num := 5
	fmt.Println("Before passValue, num:", num)
	passValue(num) // num remains unchanged
	fmt.Println("After passValue, num:", num)

	// Function with pointer parameter
	numPtr := &num
	fmt.Println("Before passByReference, num:", *numPtr)
	passByReference(numPtr) // num is modified through the pointer
	fmt.Println("After passByReference, num:", *numPtr)

	// Function with multiple return values
	message, err := multiReturn(true)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Message:", message)
	}
}
