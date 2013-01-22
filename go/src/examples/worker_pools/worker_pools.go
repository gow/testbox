package worker_pools

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker: ", id, "Processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func worker_pools(num_worker int, num_jobs int, jobs chan int, results chan int) {
	for w := 1; w <= num_worker; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= num_jobs; j++ {
		jobs <- j
	}

	for a := 1; a <= num_jobs; a++ {
		<-results
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	worker_pools(3, 9, jobs, results)
	close(jobs)
	close(results)
}
