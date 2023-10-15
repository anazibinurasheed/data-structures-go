package main

const (
	ArraySize = 7
)

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *buckeNode
}

type buckeNode struct {
	key   string
	value string
	next  *buckeNode
}

func NewHashMap() *HashTable {

	hashTable := &HashTable{}

	for i := range hashTable.array {
		hashTable.array[i] = &bucket{}
	}
	return hashTable
}

func (h *HashTable) Insert(key, val string) {
	index := hash(key)
	h.array[index].insert(key, val)
}

func (h *bucket) insert(k, v string) {

	current := h.head

	for current != nil {
		if current.key == k {
			current.value = v
			return
		}
		current = current.next

	}

	n := &buckeNode{key: k, value: v}

	n.next = current
	current = n

}

func (h *HashTable) Delete(key, val string) {

}

func (h *bucket) delete(key, val string) {

}

func (h *HashTable) Search(key, val string) {

}

func (h *bucket) search(key, val string) {

}

func hash(key string) int {
	sum := 0

	for _, val := range key {
		sum += int(val)
	}
	return sum % ArraySize
}

func main() {



}
