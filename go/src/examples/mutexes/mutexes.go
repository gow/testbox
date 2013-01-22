package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var ops uint64 = 0

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			{
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	for w := 0; w <= 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(1000)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("State:", state)
	mutex.Unlock()
}
