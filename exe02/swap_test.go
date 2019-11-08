package main

import "testing"

func Test_Swap0test(t *testing.T) {
	a := 1
	b := 2
	c, d := swap0(a, b)
	if c != 2 || d != 1 {
		t.Error("Fail")
	}
}

func Test_Swap1test(t *testing.T) {
	a := 1
	b := 2
	swap1(a, b)
	if a != 2 || b != 1 {
		t.Error("Fail")
	}
}

func Test_swap2test(t *testing.T) {
	a := 1
	b := 2
	swap2(&a, &b)
	if a != 2 || b != 1 {
		t.Error("Fail")
	}
}

func Test_swap3test(t *testing.T) {
	a := 1
	b := 2
	pa, pb := &a, &b
	swap3(&pa, &pb)
	if a != 2 || b != 1 {
		t.Error("Fail")
	}
}
