package errors1

import "fmt"

func ExampleErrors() {
	for _, i := range []int{10, 42} {
		if r, e := fun1(i); e != nil {
			fmt.Println("Fun1 failed:", e)
		} else {
			fmt.Println("fun worked!:", r)
		}
	}

	for _, i := range []int{12, 42} {
		if r, e := fun2(i); e != nil {
			fmt.Println("Fun2 failed:", e)
		} else {
			fmt.Println("fun2 worked!", r)
		}
	}

	_, e := fun2(42)
	if ae, ok := e.(argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
	//Output:
	//fun worked!: 13
	//Fun1 failed: Can't work with 42
	//fun2 worked! 15
	//Fun2 failed: 42 - Can't work with it
	//42
	//Can't work with it

}
