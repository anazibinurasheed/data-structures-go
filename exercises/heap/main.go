package main

import "fmt"

func main() {
	heap := NewMaxHeap()

	for i := 0; i < 10; i++ {
		heap.Insert(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(heap.Extract())
	}

}

type maxHeap struct {
	array []int
}

func NewMaxHeap() *maxHeap {
	return &maxHeap{}
}

func (m *maxHeap) heapifyDown() {
	root := 0
	lastIndex := len(m.array) - 1
	l, r := left(root), right(root)
	var childToCompare int

	for l <= lastIndex {
		switch {
		case l == lastIndex:
			childToCompare = l

		case m.array[l] > m.array[r]:
			childToCompare = l

		default:
			childToCompare = r
		}

		if m.array[root] < m.array[childToCompare] {
			m.swap(root, childToCompare)

			root = childToCompare
			l, r = left(root), right(root)
			continue
		}

		return
	}

}

func (m *maxHeap) heapifyUp() {
	index := len(m.array) - 1
	for m.array[index] > m.array[parent(index)] {

		m.swap(index, parent(index))
		index = parent(index)
	}
}

func (m *maxHeap) swap(i1, i2 int) {
	m.array[i1], m.array[i2] = m.array[i2], m.array[i1]
}

func (m *maxHeap) Insert(val int) {
	m.array = append(m.array, val)
	m.heapifyUp()

}

func (m *maxHeap) Extract() int {
	if len(m.array) < 1 {
		fmt.Println("empty heap")
		return -1
	}

	extracted := m.array[0]

	l := len(m.array) - 1
	m.array[0] = m.array[l]
	m.array = m.array[:l]
	m.heapifyDown()
	return extracted

}

func (m *maxHeap) Max() int {
	if len(m.array) < 1 {
		fmt.Println("empty heap")
		return -1
	}

	return m.array[0]
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
