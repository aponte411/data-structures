package autocomplete_system

import (
	"github.com/aponte411/data-structures/non_linear/trie"
	"sort"
)

type AutoCompleteSystem struct {
	dict map[string]int
	str  string
	tr   *trie.Trie
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
		tr:   tr,
	}
}

type Result struct {
	count int
	word  string
}

func (a *AutoCompleteSystem) Input(ch string) []string {
	if ch != "#" {
		a.str += ch
	} else {
		a.dict[a.str] += 1
		a.tr.Insert(a.str)
		// Reset
		a.str = ""
		return []string{}
	}
	// Get top words with user input
	topWords := a.tr.ReturnTopWordsWithPrefix(a.str)
	// Convert to slice of Results
	results := make([]Result, 0)
	for _, word := range topWords {
		count := a.dict[word]
		res := Result{count, word}
		results = append(results, res)
	}
	// Sort by counts but if words have same count, sort by word
	sort.Slice(results, func(i, j int) bool {
		if results[i].count > results[j].count {
			return true
		} else if results[i].count < results[j].count {
			return false
		}
		return results[i].word < results[j].word
	})
	// Return top 3 words
	top3 := make([]string, 0)
    for i := range results[:3] {
		top3 = append(top3, results[i].word)
	}
	return top3
}
