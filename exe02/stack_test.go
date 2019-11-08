package main

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {

	s := NewStack()

	s.Push("1")
	s.Push("2")
	s.Push("3")

	if s.Pop() != "3" {
		t.Error("Pop() did not return 3")
	}

	if s.Pop() != "2" {
		t.Error("Pop() did not return 2")
	}

	if s.Pop() != "1" {
		t.Error("Pop() did not return 1")
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		}
	}()

	if s.Pop() != nil {
		t.Error("Stack should be empty, but is not.")
	}

	r1 := NewRational(1, 2)
	r2 := NewRational(2, 4)

	s.Push(r1)
	s.Push(r2)

	if s.Pop() != r2 {
		t.Error("Pop() did not return r2")
	}

}

func TestCasting(t *testing.T) {

	s := NewStack()
	s.Push(1)
	s.Push(2)

	sum := 0
	for i := 0; i < s.Size(); i++ {
		sum += s.Get(i).(int) // type assertion = cast from interface{} to int
	}

	if sum != 3 {
		t.Error(fmt.Sprintf("Sum result should be 3 but is %v", sum))
	}
}