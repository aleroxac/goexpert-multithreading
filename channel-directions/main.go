package main

import "fmt"

func pub(name string, hello chan<- string) { //send-only
	hello <- name
}

func sub(data <-chan string) { // receive-only
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go pub("Hello", hello)
	sub(hello)
}
