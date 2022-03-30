package main

import "fmt"

func main() {
	// buffered
	Buffering := make(chan string, 2)
	Buffering <- "buffered"
	Buffering <- "channel"

	fmt.Println(<-Buffering)
	fmt.Println(<-Buffering)
}
