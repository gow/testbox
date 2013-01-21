package methods

import "fmt"
import "testing"

func TestMethods_1(t *testing.T) {
	input := rect{3, 4}
	expected := rect{5, 6}
	x := input
	x.mutate(expected.width, expected.height)
	if x != expected {
		t.Errorf("Input: %v, output: %v, Expected: %v", input, x, expected)
	}
}

func ExampleMethods() {
	r := rect{width: 10, height: 20}
	fmt.Printf("r: [%v], area: [%d], Perimeter: [%d]\n\n", r, r.area(), r.perimeter())

	rp := &r
	fmt.Printf("rp: [%v], area: [%d], Perimeter: [%d]\n\n", rp, rp.area(), rp.perimeter())

	rp.mutate(3, 4)
	fmt.Printf("r: [%v], area: [%d], Perimeter: [%d]\n", r, r.area(), r.perimeter())
	fmt.Printf("rp: [%v], area: [%d], Perimeter: [%d]\n\n", rp, rp.area(), rp.perimeter())

	rp_dummy := &r
	r.dummyMutate(5, 6)
	rp.dummyMutate(5, 6)
	fmt.Printf("r: [%v], area: [%d], Perimeter: [%d]\n", r, r.area(), r.perimeter())
	fmt.Printf("rp_dummy: [%v], area: [%d], Perimeter: [%d]\n\n", rp_dummy, rp_dummy.area(), rp_dummy.perimeter())

	r.mutate(5, 6)
	fmt.Printf("r.mutate(5,6): [%v], area: [%d], Perimeter: [%d]\n", r, r.area(), r.perimeter())
	//Output:
	//r: [{10 20}], area: [200], Perimeter: [60]
	//
	//rp: [&{10 20}], area: [200], Perimeter: [60]
	//
	//r: [{3 4}], area: [12], Perimeter: [14]
	//rp: [&{3 4}], area: [12], Perimeter: [14]
	//
	//r: [{3 4}], area: [12], Perimeter: [14]
	//rp_dummy: [&{3 4}], area: [12], Perimeter: [14]
	//
	//r.mutate(5,6): [{5 6}], area: [30], Perimeter: [22]

}
