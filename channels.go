package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	printer := make(chan string)

	go func() {
		msg := <-messages
		printer <- msg
	}()

	go func() { messages <- "ping" }()

	output := <-printer
	fmt.Println(output)
}
