package main

// Stack represents generic stack
type Stack[T any] struct {
	values []T
}

// Push pushes values on top of stack
func (s *Stack[T]) Push(v ...T) {
	s.values = append(s.values, v...)
}

// Pop pops from top of stack
func (s *Stack[T]) Pop() T {
	i := len(s.values) - 1
	v := s.values[i]
	s.values = s.values[:i]
	return v
}

// Top gets from top of stack
func (s *Stack[T]) Top() T {
	return s.values[len(s.values)-1]
}

// Size returns size of stack
func (s *Stack[T]) Size() int {
	return len(s.values)
}

// Empty returns true if stack is empty
func (s *Stack[T]) Empty() bool {
	return len(s.values) == 0
}

// Values returns slice of values stored in stack
func (s *Stack[T]) Values() []T {
	return s.values
}
