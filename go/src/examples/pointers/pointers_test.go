package main

import "testing"

func TestZeroPtr_1(t *testing.T) {
	const in, out = 32, 0
	x := in
	zeroPtr(&x)
	if x != out {
		t.Errorf("in: %v, output: %v, Expected %v", in, x, out)
	}
}
