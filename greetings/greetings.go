package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	/*
	 * In Go, the := operator is a shortcut for declaring and initializing a variable in one line
	 * (Go uses the value on the right to determine the variable's type). Taking the long way,
	 * you might have written this as:
	 *
	 * var message string
	 * message = fmt.Sprintf(randomFormat(), name)
	 */

	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function.
// Go executes init functions automatically at program startup, after global variables have been initialized.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
// lowercased functions are accessible only to code in its their package (in other words, they are not exported).
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
