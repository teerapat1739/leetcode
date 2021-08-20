package main

import "fmt"

// AlphabetSize is the constant alpha value for the Tire Algorithm
const AlphabetSize = 26

// Node represent a node in the Tire Algorithm
type Node struct {
	child [AlphabetSize]*Node
	isEnd bool
}

// Trie represent a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie will create new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert will take a word and add it to the trie
func (t *Trie) Insert(word string) {
	wordLength := len(word)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.child[charIndex] == nil {
			currentNode.child[charIndex] = &Node{}
		}
		currentNode = currentNode.child[charIndex]
	}
	currentNode.isEnd = true
}

// Search will take a word and RETURN true is that word is included in the trie
func (t *Trie) Search(word string) bool {
	wordLength := len(word)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.child[charIndex] == nil {
			return false
		}
		currentNode = currentNode.child[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	myTrie := InitTrie()

	toAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}


	for _, v := range toAdd {
		myTrie.Insert(v)
	}

	fmt.Println(myTrie.Search("argon"))
}
