package channels

import (
	"fmt"
	"time"
)

/*
 * in: receive-only channel
 * out: send-only channel
 */
func goFun(t time.Duration, name string, in <-chan string, out chan<- string) {
	time.Sleep(t)
	msg := <-in
	fmt.Println("goroutine: ", name, "Received: ", msg)
	out <- msg
}

/*
 * in: receive-only channel
 * out: send-only channel
 */
func goFun1(t time.Duration, in <-chan string, out chan<- string) {
	name := fmt.Sprintf("goFun1 (%v)", t)
	goFun(t, name, in, out)
}

/*
 * in: receive-only channel
 * out: send-only channel
 */
func goFun2(t time.Duration, in <-chan string, out chan<- string) {
	name := fmt.Sprintf("goFun2 (%v)", t)
	goFun(t, name, in, out)
}

/*
 * in: receive-only channel
 * out: send-only channel
 */
func goFun3(t time.Duration, in <-chan string, out chan<- string) {
	name := fmt.Sprintf("goFun3 (%v)", t)
	goFun(t, name, in, out)
}

// This is intentionally left here to demonstrate test coverage.
func main() {
	chan1 := make(chan string, 2) // This channel buffers upto 2 messages.
	chan2 := make(chan string)
	chan3 := make(chan string)

	fmt.Println("Sending message to chan1...")
	chan1 <- "Test msg"
	go goFun1(2*time.Second, chan1, chan2)
	go goFun2(6*time.Second, chan2, chan3)
	go goFun3(3*time.Second, chan3, chan1)

	time.Sleep(3 * time.Second)
	msg := <-chan1
	fmt.Println("Received message from chan1: ", msg)
}
