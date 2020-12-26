package trie

// Use cases: autocomplete, search, suggestions, etc.
type TrieNode struct {
	children map[byte]*TrieNode
	word     bool
}

type Trie struct {
	root *TrieNode
}

func InitTrie() *Trie {
	root := &TrieNode{
		children: make(map[byte]*TrieNode),
		word:     false,
	}
	return &Trie{
		root: root,
	}
}

// O(N) time, O(N) space?
func (t *Trie) Insert(word string) {
	node := t.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, ok := node.children[char]; !ok {
			node.children[char] = &TrieNode{
				children: make(map[byte]*TrieNode),
				word:     false,
			}
		}
		node = node.children[char]
	}
	node.word = true
}

// O(N) time, O(1) space
func (t *Trie) Search(word string) bool {
	node := t.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, ok := node.children[char]; !ok {
			return false
		}
		node = node.children[char]
	}
	return node.word
}

// O(N) time, O(1) space
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		if _, ok := node.children[char]; !ok {
			return false
		}
		node = node.children[char]
	}
	return true
}

// Return TopWords
func (t *Trie) ReturnTopWordsWithPrefix(prefix string) []string {
	node := t.root
	for i := 0; i < len(prefix); i++ {
		ch := prefix[i]
		if _, ok := node.children[ch]; !ok {
			return []string{}
		}
		node = node.children[ch]
	}
	topWords := make([]string, 0)
	TrieDFS(node, prefix, &topWords)
	return topWords
}

// DFS
func TrieDFS(node *TrieNode, prefix string, topWords *[]string) {
	if node.word {
		*topWords = append(*topWords, prefix)
	}
	for ch := range node.children {
		// Convert prefix to slice of bytes and append current child
		// character to slice. Then convert back to string
		newPrefix := string(append([]byte(prefix), ch))
		// Recurse through child node and add child to prefix
		TrieDFS(node.children[ch], newPrefix, topWords)
	}
}
