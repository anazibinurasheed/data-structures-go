package main

import (
	"fmt"
)

/*
heap is a tree based data structure created using array. each index is mapped as child or parent.

parent = i/ 2
left = 2*i +1
right = 2*i +2

property of maxheap : the child should be less than the parent
heap is used to create priority queue

getting the high priority value from the heap will take O(1) complexity.

insertion, Deletion, Heapify : 0(logN)
*/
type maxHeap struct {
	array []int
}

func NewMaxHeap() *maxHeap {

	return &maxHeap{}
}

func (h *maxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

// heapifyUp will maintain the property of the heap when an element inserted into it.
func (h *maxHeap) heapifyUp(index int) {

	// finally when it reaches the top the parent become 0 and we try to get the parent of 0, the computation = 0/2 == 0, so no worries.
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *maxHeap) Extract() int {
	if len(h.array) == 0 {
		fmt.Println("heap is empty")
		return -1
	}

	extracted := h.array[0]

	l := len(h.array) - 1
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.HeapifyDown(0)
	return extracted

}

// HeapifyDown will heapify top to bottom
func (h *maxHeap) HeapifyDown(index int) {

	lastIndex := len(h.array) - 1 // took the last index

	l, r := left(index), right(index) // took the child of the index

	childToCompare := 0 // initialized a variable for holding index

	// loop while index has at least one child
	//The loop continues as long as the left child exists, which is ensured by l <= lastIndex. If l exceeds lastIndex, it means that the current element doesn't have a left child, and there's no need to continue the loop.

	// If we have the left child then only we want to run the loop else all is done. If the parent has any child the first child should be left because its a complete binary tree, if the parent has no left child then no need of check the loop will exit because the calculation of the child will be over than then the last index of the array. A complete binary tree is insertion technique from left child to right and the parent of the child must be greater than the child.
	for l <= lastIndex {
		switch {

		case l == lastIndex: // when left child is the only child
			childToCompare = l

		case h.array[l] > h.array[r]: // when left child is larger
			childToCompare = l

		default: // when right child is larger
			childToCompare = r
		}

		// compare array value of current index to larger child and swap if smaller

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}

}

func parent(i int) int {
	return i / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *maxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {

	heap := NewMaxHeap()

	for i := 0; ; {
		heap.Insert(i)
		i++

		if i == 10 {
			break
		}
	}

	for i := len(heap.array) - 1; i > 0; i-- {
		fmt.Println(heap.Extract())

	}

}
