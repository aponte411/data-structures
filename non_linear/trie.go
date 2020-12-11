package main

import "fmt"

// Use cases: autocomplete, search, suggestions, etc.
type TrieNode struct {
	Children map[string]*TrieNode
	end      bool
}

type Trie struct {
	Root TrieNode
}

// O(N) time, O(N) space?
func (t *Trie) Insert(word string) {
	node := t.Root
	for char := range word {
		ch := string(char)
		if _, ok := node.Children[ch]; !ok {
            node.Children[ch] = &TrieNode{end: false}
		}
		node = node.Children[ch]
	}
	node.end = true
}

// O(N) time, O(1) space
func (t *Trie) Search(word string) bool {
	root := t.Root
	for char := range word {
		ch := string(char)
		if _, ok := root.Children[ch]; !ok {
			return false
		}
		root = root.Children[ch]
	}
	return root.end
}

// O(N) time, O(1) space
func (t *Trie) StartsWith(prefix string) bool {
	root := t.Root
	for char := range prefix {
		ch := string(char)
		if _, ok := root.Children[ch]; !ok {
			return false
		}
		root = root.Children[ch]
	}
	return true
}

func main() {
	trie := new(Trie)
	word := "David"
	trie.Insert(word)
	fmt.Println(trie.Root.Children)

}
