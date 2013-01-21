package selects

func ExampleSelect() {
	var terminate int
	terminate = 3
	demonstrateSelect(2, 1, 4, terminate)

	terminate = 5
	demonstrateSelect(2, 1, 4, terminate)

	terminate = 7
	// this will terminate before 6 (because if time.After(3s) limit)
	demonstrateSelect(2, 1, 6, terminate)

	//Output:
	//Channel 2 slept for 1s
	//Channel 1 slept for 2s
	//Termination message after 3s
	//Channel 2 slept for 1s
	//Channel 1 slept for 2s
	//Channel 3 slept for 4s
	//Termination message after 5s
	//Channel 2 slept for 1s
	//Channel 1 slept for 2s
	//Termination message after 3s
}
