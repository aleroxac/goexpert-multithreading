package main

import (
	"fmt"
)

func publisher(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int) {
	for i := range ch {
		fmt.Printf("Received: %d\n", i)
	}
}

func main() {
	ch := make(chan int)
	go publisher(ch)
	consumer(ch)
}
