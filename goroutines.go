package main

import (
	"fmt"
	"time"
)

func f(thread string) {
	for i := 0; i < 3; i++ {
		fmt.Println(thread, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("goroutine 2")

	go func(msg string) {
		fmt.Println(msg)
	}("goroutine 3")

	time.Sleep(time.Second)
	fmt.Println("done")
}
