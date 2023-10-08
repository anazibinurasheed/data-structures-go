package main

import (
	"errors"
	"fmt"
)

//Heapify: a process of creating a heap from an array.

//Insertion: process to insert an element in existing heap time complexity O(log N).

//Deletion: deleting the top element of the heap or the highest priority element, and then organizing the heap and returning the element with time complexity O(log N).

// Peek: to check or find the first (or can say the top) element of the heap.

//Max-Heap: In a Max-Heap the key present at the root node must be greatest among the keys present at all of it’s children. The same property must be recursively true for all sub-trees in that Binary Tree.

//Min-Heap: In a Min-Heap the key present at the root node must be minimum among the keys present at all of it’s children. The same property must be recursively true for all sub-trees in that Binary Tree.

/*
Properties of Heap:

Heap has the following Properties:

Complete Binary Tree: A heap tree is a complete binary tree, meaning all levels of the tree are fully filled except possibly the last level, which is filled from left to right. This property ensures that the tree is efficiently represented using an array.

Heap Property: This property ensures that the minimum (or maximum) element is always at the root of the tree according to the heap type.

Parent-Child Relationship: The relationship between a parent node at index ‘i’ and its children is given by the formulas: left child at index 2i+1 and right child at index 2i+2 for 0-based indexing of node numbers.

Efficient Insertion and Removal: Insertion and removal operations in heap trees are efficient. New elements are inserted at the next available position in the bottom-rightmost level, and the heap property is restored by comparing the element with its parent and swapping if necessary. Removal of the root element involves replacing it with the last element and heapifying down.

Efficient Access to Extremal Elements: The minimum or maximum element is always at the root of the heap, allowing constant-time access.

*/
//Heapify:
//It is the process to rearrange the elements to maintain the property of heap data structure. It is done when a certain node creates an imbalance in the heap due to some operations on that node. It takes O(log N) to balance the tree.

//Insertion:
//If we insert a new element into the heap since we are adding a new element into the heap so it will distort the properties of the heap so we need to perform the heapify operation so that it maintains the property of the heap.
//This operation also takes O(logN) time.

//Deletion:

//If we delete the element from the heap it always deletes the root element of the tree and replaces it with the last element of the tree.

//Since we delete the root element from the heap it will distort the properties of the heap so we need to perform heapify operations so that it maintains the property of the heap.

// It takes O(logN) time.

// usecase :
// Priority queue, heap sort

/*
Advantages of Heaps:
Fast access to maximum/minimum element (O(1))
Efficient Insertion and Deletion operations (O(log n))
Flexible size
Can be efficiently implemented as an array
Suitable for real-time applications
Disadvantages of Heaps:
Not suitable for searching for an element other than maximum/minimum (O(n) in worst case)
Extra memory overhead to maintain heap structure
Slower than other data structures like arrays and linked lists for non-priority queue operations.
*/


type MaxHeap struct {
	arr []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) HeapifyUp(index int) {
	parent := h.parent(index)

	for index > 0 && h.arr[index] > h.arr[parent] {
		h.Swap(index, parent)
		index = parent
		parent = h.parent(index)
	}
}

func (h *MaxHeap) HeapifyDown(index int) {
	for {
		leftChild := h.lChild(index)
		rightChild := h.rChild(index)
		largest := index

		if leftChild < h.len() && h.arr[leftChild] > h.arr[largest] {
			largest = leftChild
		}

		if rightChild < h.len() && h.arr[rightChild] > h.arr[largest] {
			largest = rightChild
		}

		if largest == index {
			break
		}

		h.Swap(index, largest)
		index = largest
	}
}

func (h *MaxHeap) Insert(val int) {
	h.arr = append(h.arr, val)
	h.HeapifyUp(h.len() - 1)
}


func (h *MaxHeap) Delete() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("heap is empty")
	}
	// swap the root (max element) with the last element
	n := h.len()-1
	h.Swap(0, n)
	max := h.arr[n]
	h.arr = h.arr[:n]

	// heapify down to maintain the max heap property
	h.HeapifyDown(0)

	return max, nil

}


func (h *MaxHeap) GetMax() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("heap is empty")
	}
	return h.arr[0], nil
}

func (h *MaxHeap) IsEmpty() bool {
	return len(h.arr) == 0
}

// returns size
func (h *MaxHeap) Size() int {
	return h.len()
}

// swap  swaps between given indexes
func (h *MaxHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

// returns length
func (h *MaxHeap) len() int {
	return len(h.arr)
}

// return the index of the left child
func (h *MaxHeap) lChild(i int) int {
	return (i*2 + 1)
}

// returns the index of the right child
func (h *MaxHeap) rChild(i int) int {
	return (i*2 + 2)
}

// returns the parent of  the i th element
func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}


func main() {

	mHeap := NewMaxHeap()

	mHeap.Insert(10)
	val, _ := mHeap.GetMax()

	fmt.Println(val)
}
