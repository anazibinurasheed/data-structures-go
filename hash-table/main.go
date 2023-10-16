package main

import "fmt"

// ArraySize is the size of the hash table array
const ArraySize = 7

// HashTable will hold an array

type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list in each slot of the array
type bucket struct {
	head *bucketNode
}

// bucketNode is a linked list node that holds the address of the next node.
type bucketNode struct {
	key  string
	next *bucketNode
}

// Init will create a bucket in each slot of the hash table.
func Init() *HashTable {
	result := &HashTable{}

	for i := range result.array {
		result.array[i] = &bucket{}

	}
	return result
}

// search will take in a key and return true if that key is stored in the hash table.
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// delete will take in a key and delete from hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)

}

// insert will take a  key, create a node with the key and insert it into the bucket.
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode

	} else {
		fmt.Println(k, "key already exists")
	}
}

func (b *bucket) delete(key string) {

	if b.head.key == key {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == key {
			previousNode.next = previousNode.next.next
			return
		}
		previousNode = previousNode.next
	}

}

// search will take in a key and return true if the bucket has the key
func (b *bucket) search(key string) bool {
	currentNode := b.head

	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func main() {

	m := Init()

	m.Insert("anas")

	m.Insert("anas")

	m.Insert("anas")




	
}
