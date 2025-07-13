package advance

import "fmt"

// CustomError is a custom error type.
type CustomError struct {
	Code int
	Msg  string
}

// To be a valid error, a type must implement the Error() method.
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Msg)
}

// RiskyOperation returns an error if something goes wrong.
func RiskyOperation(shouldFail bool) (string, error) {
	if shouldFail {
		return "", &CustomError{Code: 500, Msg: "Something went wrong"}
	}
	return "Success", nil
}

// DemonstrateErrors shows Go;s idiomatic error handling
func DemonstrateErrors() {
	fmt.Println("\n[Error Handling]")

	// Good case
	if result, err := RiskyOperation(false); err != nil {
		fmt.Printf("Failed: %v\n", err)
	} else {
		fmt.Printf("Succeeded: %s\n", result)
	}

	// Bad case
	if _, err := RiskyOperation(true); err != nil {
		fmt.Printf("Failed as expected: %v\n", err)
	}

}
