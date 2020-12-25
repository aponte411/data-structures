package linear

import "testing"

func TestSLList(t *testing.T){

	list := new(SinglyLinkedList)
	list.AddToHead(1)
	list.AddToHead(2)
	list.AddToEnd(4)
	list.AddToEnd(77)
	list.AddAfter(1, 9)
	fmt.Printf("Count: %v\n", list.count)
	list.Iterate()
	list.Reverse()
	list.Iterate()
	list.Reverse()
	list.Iterate()
	// reverse print
	list.ReversePrint(list.Head)

}
