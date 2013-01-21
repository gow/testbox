package selects

import (
	"fmt"
	"time"
)

// This function creates a couple of channels and a terminating channel. Different goroutines write
// messages to these channels after a sleep time. These messages are selected in an infinite loop.
// The function stops selecting when ther terminating channel delivers the message
func demonstrateSelect(t1 int, t2 int, t3 int, term int) {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	tc := make(chan string)

	go func(sec int, c chan<- string) {
		t := time.Duration(sec) * time.Second
		time.Sleep(t)
		c <- fmt.Sprintf("Channel 1 slept for %v", t)
	}(t1, c1)

	go func(sec int, c chan<- string) {
		t := time.Duration(sec) * time.Second
		time.Sleep(t)
		c <- fmt.Sprintf("Channel 2 slept for %v", t)
	}(t2, c2)

	go func(sec int, c chan<- string) {
		t := time.Duration(sec) * time.Second
		time.Sleep(t)
		c <- fmt.Sprintf("Channel 3 slept for %v", t)
	}(t3, c3)

	go func(sec int, c chan<- string) {
		t := time.Duration(sec) * time.Second
		time.Sleep(t)
		c <- fmt.Sprintf("Termination message after %v", t)
	}(term, tc)

infinite_for:
	for {
		select {
		// This starts off a timer for 3 secs. Ensures that a single iteration doesn't run
		// for more than 3 secs
		case <-time.After(3 * time.Second):
			fmt.Printf("Termination message after 3s\n")
			break infinite_for // break the infinite for loop
		case msg := <-tc:
			fmt.Println(msg)
			break infinite_for // break the infinite for loop
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c3:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}
	}
}

func main() {
	demonstrateSelect(2, 4, 6, 5)
}
