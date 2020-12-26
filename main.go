package main

import (
	"fmt"

	"github.com/aponte411/data-structures/non_linear/trie"
)

func main() {
	tr := trie.InitTrie()
	tr.Insert("paticakes")
	fmt.Println(tr)
}
