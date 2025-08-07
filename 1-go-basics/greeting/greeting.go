// Package greeting provides functions to return greetings
package greeting

import "fmt"

// SayHello returns a greeting message
func SayHello(name string) string{
	return fmt.Sprintf("Hello, %s! Welcome to Go packages.", name)
}

// SayGoodbye return a farewell message
func SayGoodbye(name string) string {
	return fmt.Sprintf("Goodbye, %s! See you next time.", name)
}