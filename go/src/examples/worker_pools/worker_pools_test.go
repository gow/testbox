package worker_pools

import "testing"

func TestWorkerPools(t *testing.T) {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	worker_pools(4, 20, jobs, results)

	close(jobs)
	close(results)
}

func ExampleWorkerPools() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	worker_pools(3, 9, jobs, results)

	close(jobs)
	close(results)

}
