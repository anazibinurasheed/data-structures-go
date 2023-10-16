package main

import "fmt"

// a trie is a tree based data structure it is mostly used for storing words

// application : auto complete
// in trie each node will have 26 children nodes for all the possible letters.
// each of the node will have bool value to indicate weather it is end of a word or not

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

// Node represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represent a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie  will create a new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {

	wordLength := len(w)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true

}

// Search will take in a word and return true is that word is included in the trie
func (t *Trie) Search(w string) bool {
	currentNode := t.root
	if currentNode == nil {
		return false
	}
	for _, v := range w {
		charIndex := v - 'a'

		if currentNode.children[charIndex] == nil {
			return false
		}

		currentNode = currentNode.children[charIndex]
	}

	return currentNode.isEnd

}

func main() {
	testTrie := InitTrie()
	testTrie.Insert("hello")

	fmt.Println(testTrie.Search("hello"))
}
