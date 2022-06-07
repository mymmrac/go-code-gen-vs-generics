package main

import "fmt"

// ======== Stack Operations =========
// Push, Pop, Top, Size, Empty, Values

// ==== Code Generation ====
//go:generate go run stack-generator.go -o stack-int.go -p main -n Int -t int

func main() {
	// ==== Generic Stack ====
	s1 := Stack[int]{}

	s1.Push(10, 20, 30)
	fmt.Println(s1.Values())

	for !s1.Empty() {
		fmt.Println(s1.Pop(), s1.Size())
	}

	s1.Push(42)
	fmt.Println(s1.Top(), s1.Empty())

	// ==== Generated Stack ====
	s2 := StackInt{}

	s2.Push(10, 20, 30)
	fmt.Println(s2.Values())

	for !s2.Empty() {
		fmt.Println(s2.Pop(), s2.Size())
	}

	s2.Push(42)
	fmt.Println(s2.Top(), s2.Empty())
}
