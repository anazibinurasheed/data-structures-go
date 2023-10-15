package main

import "fmt"

// Node represents the components of a binary search tree
type Node struct {
	key   int
	left  *Node
	right *Node
}

// insert will add a node to the tree
// the bst not hold any duplicate values
func (n *Node) Insert(key int) {
	if n.key < key {
		// move right
		if n.right == nil {
			n.right = &Node{key: key}

		} else {
			// we are calling the insert method on the right node
			n.right.Insert(key)
		}
	} else if n.key > key {
		// move left
		if n.left == nil {
			n.left = &Node{key: key}
			return
		} else {

			// we are calling the insert method on the left node
			n.left.Insert(key)
		}
	}

}

// search takes a key as input and return true if the key is existing in the tree.
func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}
	if n.key > key {
		// move left
		return n.left.Search(key)
	} else if n.key <key {
		// move right
		return n.right.Search(key)
	}
	return true
}

func main() {
	t := &Node{key: 100}

	t.Insert(55)
	t.Insert(66)
	t.Insert(111)
	t.Insert(124)
	t.Insert(135)
	t.Insert(77)
	fmt.Println(t)


	fmt.Println(t.Search(77))

}
