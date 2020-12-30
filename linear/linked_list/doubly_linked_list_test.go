package linked_list

import (
	"fmt"
	"testing"
)

func TestDLList(t *testing.T) {
	list := new(DoublyLinkedList)
	list.AddToHead("David")
	list.AddToHead("Pati")
	root := list.Head
	for root != nil {
		fmt.Println(root.Val)
		root = root.Next
	}
}
