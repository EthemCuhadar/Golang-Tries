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
	fmt.Printf("%s is inserted to trie\n", w)
}

// Search - Return true if word is included in.
func (t *Trie) Search(w string)bool{
	wordLength := len(w)
	currentNode := t.root
	for i:=0; i<wordLength; i++{
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil{
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true{
		fmt.Printf("Gotcha!, %s is included in the trie\n", w)
		return true
	}
	return false
}

func main() {
	countryTrie := InitTrie()
	
	toAdd := []string{
		"england",
		"france",
		"germany",
		"netherlands",
		"sweden",
		"norway",
		"spain",
		"denmark",
		"turkey",
		"greece",
	}
	
	for _, v := range toAdd{
		countryTrie.Insert(v)
	}
	fmt.Println(countryTrie.Search("germany"))
}
