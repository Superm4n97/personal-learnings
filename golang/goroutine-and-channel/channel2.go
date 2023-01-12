package main

import (
	"fmt"
	"time"
)

/*
jobs <- chan int {receiver channel}
jobs chan <- int {sender channel}
*/

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs { // we can iterate a channel, until the channel is open
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2

		// can not use this. because it's a sender channel function for 'results', and receiver channel function for 'jobs'
		//fmt.Println(<-results)
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// a channel can be closed, after closing the channel the workers loop will be break
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		<-results
	}
}
