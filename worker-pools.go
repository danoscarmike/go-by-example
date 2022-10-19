package main

import (
	"fmt"
	"time"
)

func pool_worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 0; w <= 3; w++ {
		go pool_worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	var res []int
	for a := 1; a <= numJobs; a++ {
		res = append(res, <-results)
	}
	fmt.Println(res)
}
