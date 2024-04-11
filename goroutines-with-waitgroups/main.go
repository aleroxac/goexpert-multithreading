package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Wait()
	}
}

func main() {
	wait_group := sync.WaitGroup{}
	wait_group.Add(25)

	go task("A", &wait_group)
	go task("B", &wait_group)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: Anonymous Task %s is running\n", i, "C")
			time.Sleep(1 * time.Second)
			wait_group.Done()
		}
	}()
	wait_group.Wait()
}
