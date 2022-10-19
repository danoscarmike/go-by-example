package main

import (
	"fmt"
	"sync"
	"time"
)

func wg_worker(id int) {
	fmt.Printf("worker %d started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			wg_worker(i)
		}()
	}

	wg.Wait()
}
