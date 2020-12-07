package main

import (
	"fmt"
	"heytobi.dev/greetings"
)

func main() {
	message := greetings.Hello("Golang")
	fmt.Println(message)
}
