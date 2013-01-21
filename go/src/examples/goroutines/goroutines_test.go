package goroutines

func ExampleGoroutines() {
	testGoRoutines()
	//Output:
	//Direct call : 0
	//Direct call : 1
	//Direct call : 2
	//Done Direct call
	//goroutine : 0
	//goroutine : 1
	//goroutine : 2
	//Done goroutine
	//go anonymous
	//Done testGoRoutines
}
