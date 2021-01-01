package linked_list

import (
	"fmt"
	"testing"
)

func TestDLList(t *testing.T) {
	list := new(DoublyLinkedList)
	list.AddToHead("David")
	list.AddToHead("Pati")
	list.AddToTail("Peter")
	root := list.Head
	for root != nil {
		fmt.Println(root.Val)
		root = root.Next
	}
	list.RemoveHead()
	head := list.Head
	if head.Val != "David" {
		t.Errorf("Expected David, got %v", head.Val)
	}
	list.RemoveTail()
	if list.Head.Val != "David" {
		t.Errorf("Expected Peter, got %v", list.Head.Val)
	}
}
