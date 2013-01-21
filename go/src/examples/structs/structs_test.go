package structs

import "fmt"
import "testing"

func TestStructs(t *testing.T) {
	s := person{"Chetan", 29}
	if s.age != 29 {
		t.Errorf("output: %v, Expected %v", s.age, 29)
	}
	s.age = 40
	if s.age != 40 {
		t.Errorf("output: %v, Expected %v", s.age, 40)
	}
}

func ExampleStructs() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})

	s := person{"Sean", 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age) //This works. No need to explicitly dereference the pointer.

	sp.age = 31
	fmt.Println(sp.age)

	//Output: {Bob 20}
	//{Alice 30}
	//{Fred 0}
	//&{Ann 40}
	//Sean
	//50
	//31

}
