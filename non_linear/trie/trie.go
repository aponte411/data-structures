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
	dfs(node, prefix, &topWords)
	return topWords
}

// DFS
func dfs(node *TrieNode, prefix string, res *[]string) {
	if node == nil {
		*res = append(*res, prefix)
	}
	for child := range node.children {
		dfs(node.children[child], prefix+string(child), res)
	}
}
