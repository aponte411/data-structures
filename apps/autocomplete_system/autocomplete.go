package autocomplete_system

import "linear/trie"

type AutoCompleteSystem struct {
	dict map[string]int
	str  string
	tr   *trie.TrieNode
}

func InitAutoCompleteSystem(sentences []string, times []int) AutoCompleteSystem {
	dict := make(map[string]int, 0)
	for i := 0; i < len(sentences); i++ {
		dict[sentences[i]] = times[i]
	}
	str := ""
	tr := trie.InitTrie()
	for i := 0; i < len(sentences); i++ {
		tr.Insert(sentences[i])
	}
	return AutoCompleteSystem{
		dict: dict,
		str:  str,
		trie: trie,
	}
}

func (a *AutoCompleteSystem) Input(c byte) []string {
	return []string{"hello"}
}
