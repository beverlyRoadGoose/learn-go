package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	/*
	 * In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on
	 * the right to determine the variable's type). Taking the long way, you might have written this as:
	 *
	 * var message string
	 * message = fmt.Sprintf("Hi, %v. Welcome!", name)
	 */

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}