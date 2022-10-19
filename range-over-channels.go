package main

import (
	"fmt"
	"time"
)

func main() {

	queue := make(chan string, 10)
	go func() {
		for j := 0; j <= 10; j++ {
			time.Sleep(time.Nanosecond)
			queue <- fmt.Sprintf("task: %d", j)
		}
		close(queue)
	}()

	for elem := range queue {
		fmt.Println(elem, len(queue))
	}
}
