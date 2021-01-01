package linked_list

import (
	"testing"
)

func TestSLList(t *testing.T) {

	list := new(SinglyLinkedList)
	list.AddToHead(1)
	list.AddToHead(2)
	list.AddToEnd(4)
	list.AddToEnd(77)
	list.AddAfter(1, 9)
	res := list.Count
	if res != 5 {
		t.Errorf("Expected 5 nodes in singly linked list, got %v", res)
	}
	list.Iterate()
	list.Reverse()
	list.Iterate()
	list.Reverse()
	list.Iterate()
	// reverse print
	list.ReversePrint(list.Head)

}
