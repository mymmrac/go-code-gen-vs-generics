// Code generated by stack-generator. DO NOT EDIT.

package main

// StackInt represents int stack
type StackInt struct {
	values []int
}

// Push pushes int values on top of stack
func (s *StackInt) Push(v ...int) {
	s.values = append(s.values, v...)
}

// Pop pops int from top of stack
func (s *StackInt) Pop() int {
	i := len(s.values) - 1
	v := s.values[i]
	s.values = s.values[:i]
	return v
}

// Top gets int from top of stack
func (s *StackInt) Top() int {
	return s.values[len(s.values)-1]
}

// Size returns size of stack
func (s *StackInt) Size() int {
	return len(s.values)
}

// Empty returns true if stack is empty
func (s *StackInt) Empty() bool {
	return len(s.values) == 0
}

// Values returns slice of int stored in stack
func (s *StackInt) Values() []int {
	return s.values
}