package main

// Stack is a generic LIFO container for untyped object.
type Stack struct {
	data []interface{} // kann nichts, ist leer
}

// NewStack constructs an empty stack.
func NewStack() *Stack {
	return new(Stack)
}

// Push pushes a value on the stack.
func (s *Stack) Push(value interface{}) {
	s.data = append(s.data, value)
}

// Pop pops a value from the stack. It returns an error if the stack is empty.
func (s *Stack) Pop() interface{} {
	if len(s.data) == 0 {
		panic("can not pop: empty stack")
	}
	var result = s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return result
}

// Size returns the count of elements in the Stack
func (s *Stack) Size() int {
	return len(s.data)
}

// Get returns the n-th element in the Stack
func (s *Stack) Get(idx int) interface{} {
	return s.data[idx]
}