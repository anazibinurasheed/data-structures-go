package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []interface{}
}

// creates a new stack
func NewStack() *Stack {
	return &Stack{}
}

// add elements to the stack
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// pop removes and returns the top item from the stack.
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil

}

// peek returns the top item from the stack without removing it.
func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

func (s *Stack) PrintStack() {
	for i, e := range s.items {
		fmt.Printf("index : %v, element : %v\n", i, e)
	}
}

func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	stack := NewStack()
	stack.Push("hello")
	stack.Push(1)
	stack.Push(true)

	stack.PrintStack()
}
