package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := InitTrie()
	word := "David"
	trie.Insert(word)
	res := trie.StartsWith("Da")
	if !res {
		t.Errorf("Should be true")
	}
}
