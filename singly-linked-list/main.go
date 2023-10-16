package main

import (
	"fmt"
)

type node struct {
	data int64
	next *node
}

type linkedList struct {
	head *node
	len  uint64
}

func NewLinkedList() *linkedList {
	return &linkedList{}
}

func (l *linkedList) Prepend(data int64) {
	newNode := &node{data: data}

	if l.head == nil {
		l.head = newNode
		l.len++
		return
	}

	newNode.next = l.head
	l.head = newNode
	l.len++
}

func (l *linkedList) Remove(data int64) {
	if l.head == nil {
		return
	}

	if l.head.data == data {
		l.head = l.head.next
		l.len--

		return
	}

	temp := l.head
	current := l.head.next

	for current != nil && current.data != data {
		current = current.next
	}

	if current != nil {
		temp.next = current.next
		l.len--
	}

}

func (l *linkedList) Traverse() {
	current := l.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
	fmt.Println()
}

func (l *linkedList) Contains(data int64) bool {
	current := l.head
	for current != nil {
		if current.data == data {
			return true
		}
		current = current.next
	}
	return false
}

func (l *linkedList) Length() uint64 {
	return l.len
}

// pass the head
func (l *linkedList) PrintReverse(head *node) {
	if head == nil {
		return
	}

	l.PrintReverse(head.next)
	fmt.Println(head.data)
}

// when data stores in a linked list it will be a sequence of nodes.
// the linked list is linked with each other.
// each node contains a address to the next node.
// the linked list is like an array but the in linked list the data's are stored in nodes not in indexes.

// if we want to delete any particular value from the linked list we need to traverse till the value matches to the value .

// in array if we an index of that element we can simply change the value.

// adding or deleting an element from the beginning of the linked list takes 0(1), in some application that may useful.

// if we want to insert an element in the beginning of an array we should need to re-allocate the array and move the elements .

// singly linked list contains the address of the next node.

// doubly linked list contains the address of the next and previous node. adding the element at end is simpler than singly linked list.
func main() {

	list := NewLinkedList()

	list.Prepend(89)
	list.Prepend(99)
	list.Traverse()
	list.Traverse()
	fmt.Println(list.Contains(99))
	fmt.Println(list.Length())
	list.PrintReverse(list.head)

}
