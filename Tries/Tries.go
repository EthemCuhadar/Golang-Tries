package main

import (
	"fmt"
)

// AlphabetSize is the max possible characters in the trie
const AlphabetSize = 26

// Node
type Node struct {
	children [26]*Node
	isEnd    bool
}

// Trie
type Trie struct {
	root *Node
}

// InitTrie - Will create new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert - Adds a word in Trie structure
func (t *Trie) Insert(w string){
	wordLength := len(w)
	currentNode := t.root
	for i:=0; i<wordLength; i++{
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil{
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
	fmt.Println("word is inserted to trie")
}

// Search

func main() {
	testTrie := InitTrie()
	fmt.Println(testTrie.root)
	testTrie.Insert("orc")
}
