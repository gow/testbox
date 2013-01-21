package interfaces

import (
	"fmt"
	"math"
	"testing"
)

func ExampleInterfaces() {
	s := square{7}
	shape, area, peri := measure(s)
	fmt.Printf("Shape: %s(%v), area: %f, perimeter %f\n", shape, s, area, peri)

	c := circle{radius: 10}
	shape, area, peri = measure(c)
	fmt.Printf("Shape: %s(%v), area: %f, perimeter %f\n", shape, c, area, peri)

	interface_arr := []geometry{circle{3}, square{4}, square{3}, circle{1}}
	for _, obj := range interface_arr {
		shape, area, peri := measure(obj)
		fmt.Printf("Shape: %s(%v), area: %f, perimeter %f\n", shape, obj, area, peri)
	}
	//Output:
	//Shape: Square({7}), area: 49.000000, perimeter 28.000000
	//Shape: Circle({10}), area: 314.159265, perimeter 62.831853
	//Shape: Circle({3}), area: 28.274334, perimeter 18.849556
	//Shape: Square({4}), area: 16.000000, perimeter 16.000000
	//Shape: Square({3}), area: 9.000000, perimeter 12.000000
	//Shape: Circle({1}), area: 3.141593, perimeter 6.283185
}

func TestCircle(t *testing.T) {
	input, expected_area, expected_peri := 3.0, 28.274334, 18.849556
	_, output_area, output_peri := measure(circle{input})
	if (output_area-28.274334) > 0.01 || (output_peri-18.849556) > 0.01 {
		t.Errorf("Input %f, output_area: %f, output_peri: %f [Expected: area: %f, peri: %f]",
			input,
			output_area,
			output_peri,
			expected_area,
			expected_peri,
		)
	}
}

func TestSquare(t *testing.T) {
	input, expected_area, expected_peri := 10.0, 100.0, 40.01
	_, out_area, out_peri := measure(square{input})
	if math.Abs(expected_area-out_area) > 0.01 || math.Abs(expected_peri-out_peri) > 0.01 {
		t.Errorf("Input %f, output_area: %f, output_peri: %f [Expected: area: %f, peri: %f]",
			input,
			out_area,
			out_peri,
			expected_area,
			expected_peri,
		)
	}
}
