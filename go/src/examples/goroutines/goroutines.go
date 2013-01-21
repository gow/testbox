package goroutines

import "fmt"
import "time"

func gofun(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println(name, ":", i)
	}
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Done", name)
}

func testGoRoutines() {
	gofun("Direct call")

	go gofun("goroutine")

	go func(msg string) {
		time.Sleep(4000 * time.Millisecond)
		fmt.Println(msg)
	}("go anonymous")

	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Done testGoRoutines")
}
