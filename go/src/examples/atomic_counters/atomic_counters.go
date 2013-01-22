package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var shared uint64 = 0

	for i := 1; i < 9999; i++ {
		go func() {
			time.Sleep(time.Millisecond)
			atomic.AddUint64(&shared, 1)
		}()
	}

	time.Sleep(time.Second)

	finalValue := atomic.LoadUint64(&shared)
	fmt.Println("final value:", finalValue)
}
