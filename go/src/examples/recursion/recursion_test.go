package main

import "testing"

func TestFactorial(t *testing.T) {
	const in, out = 1, 1
	if x := factorial(in); x != out {
		t.Errorf("Factorial(%v) = %v, Expected %v", in, x, out)
	}
}

func TestFactorial1(t *testing.T) {
	const in, out = 5, 120
	if x := factorial(in); x != out {
		t.Errorf("Factorial(%v) = %v, Expected %v", in, x, out)
	}
}
