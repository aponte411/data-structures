package non_linear

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := InitTrie()
	word := "David"
	trie.Insert(word)
	fmt.Println(trie.StartsWith("Da"))
}
