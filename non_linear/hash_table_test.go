package non_linear

import (
	"fmt"
	"testing"
)

func TestHashTable(t *testing.T) {
	//hashtable := &HashTable{} // or new(HashTable)
	hashtable := new(HashTable)
	hashtable.Init()
	hashtable.Add(1, 10)
	hashtable.Add(2, 20)
	hashtable.Add(3, 30)
	hashtable.Print()
	fmt.Println(hashtable.Get(1))
}
