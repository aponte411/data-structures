package hash_table

import (
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

	res := hashtable.Get(1)
	if res != 10 {
		t.Errorf("Expected 10 got %v", res)
	}
}
