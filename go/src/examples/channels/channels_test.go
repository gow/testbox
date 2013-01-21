package channels

import (
	"fmt"
	"testing"
	"time"
)

func TestChannels(t *testing.T) {
	chan1 := make(chan string, 3) // chan1 buffers upto 3 messages.
	chan2 := make(chan string)
	chan3 := make(chan string)

	const in = "[Test message]"
	const expected string = in
	chan1 <- "junk"
	chan1 <- in
	go goFun1(1*time.Second, chan1, chan2)
	go goFun2(3*time.Second, chan2, chan3)
	go goFun3(5*time.Second, chan3, chan1)

	<-chan1 // throw away "junk"
	if out := <-chan1; expected != out {
		t.Errorf("Input: [%v], output: [%v], Expected: [%v]", in, out, expected)
	}
}

func ExampleChannels() {
	chan1 := make(chan string, 1) // this channel buffers upto one message.
	chan2 := make(chan string)
	chan3 := make(chan string)

	fmt.Println("Sending message to chan1...")
	chan1 <- "Test msg"
	go goFun1(1*time.Second, chan1, chan2)
	go goFun2(3*time.Second, chan2, chan3)
	go goFun3(5*time.Second, chan3, chan1)

	time.Sleep(2 * time.Second)
	msg := <-chan1
	fmt.Println("Received message from chan1: ", msg)
	//OUtput:
	//Sending message to chan1...
	//goroutine:  goFun1 (1s) Received:  Test msg
	//goroutine:  goFun2 (3s) Received:  Test msg
	//goroutine:  goFun3 (5s) Received:  Test msg
	//Received message from chan1:  Test msg

}
