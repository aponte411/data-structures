package main

import "fmt"

// Use cases: autocomplete, search, suggestions, etc.
type TrieNode struct {
	children map[string]*TrieNode
	end      bool
}

type Trie struct {
	Root *TrieNode
}

func InitTrie() *Trie {
	root := &TrieNode{
		children: map[string]*TrieNode{},
		end:      false,
	}
	return &Trie{
		Root: root,
	}
}

// O(N) time, O(N) space?
func (t *Trie) Insert(word string) {
	node := t.Root
	for _, char := range word {
		ch := string(char)
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &TrieNode{}
		}
		node = node.children[ch]
	}
	node.end = true
}

// O(N) time, O(1) space
func (t *Trie) Search(word string) bool {
	node := t.Root
	for _, char := range word {
		ch := string(char)
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return node.end
}

// O(N) time, O(1) space
func (t *Trie) StartsWith(prefix string) bool {
	node := t.Root
	for _, char := range prefix {
		ch := string(char)
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return true
}

func main() {
	trie := InitTrie()
	word := "David"
	trie.Insert(word)
	fmt.Println(trie.Root.children)
}
