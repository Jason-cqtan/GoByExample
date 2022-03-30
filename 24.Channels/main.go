package main

import "fmt"

func main() {

	messages := make(chan string)

	// send channel <-
	go func() { messages <- "ping" }()

	// receives  <-channel
	msg := <-messages
	fmt.Println(msg)

	// buffered
	Buffering := make(chan string, 2)
	Buffering <- "buffered"
	Buffering <- "channel"

	fmt.Println(<-Buffering)
	fmt.Println(<-Buffering)
}
