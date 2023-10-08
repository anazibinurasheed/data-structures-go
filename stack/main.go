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
	for _, e := range s.items {
		fmt.Printf(" element : %v\n", e)
	}
}

func (s *Stack) Size() int {
	return len(s.items)
}

//: Stacks are very efficient for adding and removing elements. This is because elements are always added and removed from the top of the stack, which is called the Last-In-First-Out (LIFO) principle.

//LIFO

// use-case

// Stacks are used to manage function calls and memory allocation in operating systems.

// Stacks are used to implement undo/redo functionality in many software applications.

// Stacks are used to implement backtracking algorithms

func Undo(undoStack, stack *Stack) {
	if stack.IsEmpty() {
		return
	}
	val, err := stack.Pop()
	if err != nil {
		return
	}
	undoStack.Push(val)
	fmt.Println("undo")
}

func Redo(stack, undoStack *Stack) {
	if undoStack.IsEmpty() {
		return
	}

	val, err := undoStack.Pop()
	if err != nil {
		return
	}
	stack.Push(val)

	fmt.Println("redo")
}

func main() {
	stack := NewStack()
	undoStack := NewStack()

	stack.Push("hello")
	stack.Push(1)
	stack.Push(true)

	Undo(undoStack, stack)
	undoStack.PrintStack()

	Redo(stack, undoStack)
	stack.PrintStack()

	fmt.Println("undo stack")
	undoStack.PrintStack()
}
