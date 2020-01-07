package main

import "fmt"

// pings is a receiver only channel
func ping(pings chan<- string, msg string) {
	// send message to pings
	pings <- msg
}

// pings is a sender only channel and pongs is a receiver only
func pong(pings <-chan string, pongs chan<- string) {
	// put pings into msg
	msg := <-pings
	// send msg to pongs
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	// channel of buffer size 2. If channel is blocked until full
	pings2 := make(chan string, 2)
	ping(pings2, "passed another message")
	ping(pings2, "final message")
	pong(pings2, pongs)
	// pongs channel is now unblocked
	fmt.Println(<-pongs)
	// pings2 still holds one more msg
	pong(pings2, pongs)
	fmt.Println(<-pongs)

}
