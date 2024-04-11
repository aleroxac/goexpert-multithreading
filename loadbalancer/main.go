package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for msg := range data {
		fmt.Printf("[worker-%d] - received %d\n", workerId, msg)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	worker_count := 1000000

	// init workers
	for i := 0; i < worker_count; i++ {
		go worker(i, data)
	}

	for i := 0; i < 1000000; i++ {
		data <- i
	}
}
