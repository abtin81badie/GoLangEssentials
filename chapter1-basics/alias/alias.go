package alias

import (
	"fmt"
	"strings"
	"time"
)

// **1. String Alias with Custom Method**
// MyCustomString is a user-defined type based on the built-in string type.
// This allows us to define methods specific to MyCustomString.
type MyCustomString string

// IsDate attempts to parse the MyCustomString as a date using the `time.DateOnly` format ("2006-01-02").
func (s MyCustomString) IsDate() (time.Time, bool) {
	t, err := time.Parse(time.DateOnly, string(s))
	return t, err == nil
}

// ToUpper converts MyCustomString to uppercase.
func (s MyCustomString) ToUpper() MyCustomString {
	return MyCustomString(strings.ToUpper(string(s)))
}

// **2. Integer Alias with Custom Method**
// MyCustomInt is an alias for the int type, allowing custom operations.
type MyCustomInt int

// IsEven checks if MyCustomInt is even.
func (n MyCustomInt) IsEven() bool {
	return n%2 == 0
}

// Factorial computes the factorial of MyCustomInt.
func (n MyCustomInt) Factorial() (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial not defined for negative numbers")
	}
	result := 1
	for i := 2; i <= int(n); i++ {
		result *= i
	}
	return result, nil
}

// **3. Function Alias**
// MyOperation is a function type alias that takes two integers and returns an integer.
type MyOperation func(int, int) int

// Add is an example function that can be used with MyOperation.
func Add(a, b int) int {
	return a + b
}

// Multiply is another function that can be used with MyOperation.
func Multiply(a, b int) int {
	return a * b
}

// ApplyOperation applies the given MyOperation function to two integers.
func ApplyOperation(a, b int, op MyOperation) int {
	return op(a, b)
}
