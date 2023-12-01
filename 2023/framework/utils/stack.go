package utils

import "fmt"

type Stack[T any] struct {
	s []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{make([]T, 0)}
}

func (s Stack[T]) Empty() bool {
	return len(s.s) == 0
}

func (s *Stack[T]) Push(elm T) {
	s.s = append(s.s, elm)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Empty() {
		var x T
		return x, fmt.Errorf("stack is empty")
	}

	l := len(s.s)
	elm := s.s[l-1]
	s.s = s.s[:l-1]
	return elm, nil
}

func (s Stack[T]) Peek() (T, error) {
	if s.Empty() {
		var x T
		return x, fmt.Errorf("stack is empty")
	}
	l := len(s.s)
	return s.s[l-1], nil
}

func TestStack(e string) int {
	s := NewStack[int]()

	if !s.Empty() {
		fmt.Println("New stack should be empty.")
		return 1
	}

	s.Push(1)
	s.Push(2)
	elm, err := s.Peek()
	if err != nil {
		fmt.Printf("Failed to peek: %v\n", err)
		return 1
	}

	if elm != 2 {
		fmt.Printf("Peek returned the wrong value: %d\n", elm)
		return 1
	}

	elm, err = s.Pop()
	if err != nil {
		fmt.Printf("Failed to pop: %v\n", err)
		return 1
	}

	if elm != 2 {
		fmt.Printf("Pop returned the wrong value: %d\n", elm)
		return 1
	}

	elm, err = s.Peek()
	if err != nil {
		fmt.Printf("Failed to peek: %v\n", err)
		return 1
	}

	if elm != 1 {
		fmt.Printf("Peek returned the wrong value: %d\n", elm)
		return 1
	}

	elm, err = s.Pop()
	if err != nil {
		fmt.Printf("Failed to pop: %v\n", err)
		return 1
	}

	if elm != 1 {
		fmt.Printf("Pop returned the wrong value: %d\n", elm)
		return 1
	}

	if !s.Empty() {
		fmt.Println("Stack should be empty after popping everything")
		return 1
	}

	elm, err = s.Pop()

	if err == nil {
		fmt.Println("Should not be able to pop from an empty stack.")
		return 1
	}
	fmt.Println("Stack: ok")
	return 0
}
