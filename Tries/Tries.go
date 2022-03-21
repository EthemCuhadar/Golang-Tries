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

// Insert

// Search

func main() {
	testTrie := InitTrie()
	fmt.Println(testTrie.root)
}
