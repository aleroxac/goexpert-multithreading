package main

import (
	"fmt"
	"sync"
)

func publisher(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Printf("Received: %d\n", i)
		wg.Done()
	}
}

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	go publisher(ch)
	go consumer(ch, &wg)
	wg.Wait()
}
