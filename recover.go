package main

import "fmt"

func mayPanic() {
	panic("Houston!")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r)
		}
	}()

	mayPanic()

	fmt.Println("This is after mayPanic() => you will never see this.")
}
