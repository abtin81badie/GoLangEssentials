// Package mathutils provides basic mathematical operations.
package mathutils

import (
	"fmt"
	"math"
)

// Add returns the sum of two numbers.
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference between two numbers.
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two numbers.
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the quotient of two numbers.
// Returns an error if division by zero is attempted.
func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("error: division by zero is not allowed")
	}
	return float64(a) / float64(b), nil
}

// Power raises a number `base` to the power of `exp`.
func Power(base, exp float64) float64 {
	return math.Pow(base, exp)
}

// SquareRoot returns the square root of a number.
// Returns an error if input is negative.
func SquareRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("error: cannot compute square root of a negative number")
	}
	return math.Sqrt(num), nil
}

// Factorial calculates the factorial of a given number (n!).
// Returns 1 if `n` is 0, and an error if `n` is negative.
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("error: factorial is not defined for negative numbers")
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}
