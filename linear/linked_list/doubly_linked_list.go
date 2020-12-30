package linked_list

// DLNode: Doubly Linked Node
type DoublyLinkedNode struct {
	Val  interface{}
	Prev *DoublyLinkedNode
	Next *DoublyLinkedNode
}

type DoublyLinkedList struct {
	Head  *DoublyLinkedNode
	Tail  *DoublyLinkedNode
	Count int
}

// NodeBetweenValues
// O(N) time, O(1) space
func (d *DoublyLinkedList) NodeBetweenValues(first, second int) *DLNode {
	var nodeWith *DLNode
	for node := d.Head; node != nil; node = node.Next {
		if node.Prev.Val == first && node.Next.Val == second {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddToHead method
// O(1) time, O(1) space
func (d *DoublyLinkedList) AddToHead(val interface{}) {
	newNode := &DLNode{val, nil, nil}
	if d.Count == 0 {
		d.Tail = newNode
		d.Head = newNode
	} else {
		d.Head.Prev = newNode
		newNode.Next = d.Head
		d.Head = newNode
	}
	d.Count += 1
}

// SortedInsert method
// O(N) time, O(1) space
func (d *DoublyLinkedList) SortedInsert(val interface{}) {
	newNode := &DLNode{val, nil, nil}
	d.Count += 1
	curr := d.Head
	// if val is first element
	if curr == nil {
		d.Tail = newNode
		d.Head = newNode
	}
	// if val is at beginning
	if curr.Val <= val {
		// make newNode head
		newNode.Next = d.Head
		d.Head.Prev = newNode
		d.Head = newNode
	}

	for curr.Next != nil && curr.Val > val {
		curr = curr.Next
	}
	// modify tail
	if curr.Next == nil {
		d.Tail = newNode
		newNode.Prev = curr
		curr.Next = newNode
	} else {
		// all other general cases
		newNode.Next = curr.Next
		newNode.Prev = curr
		curr.Next = newNode
		newNode.Next.Prev = newNode
	}
}

// RemoveHead
// O(1) time, O(1) space
func (d *DoublyLinkedList) RemoveHead() bool {
	if d.count == 0 {
		return false
	}
	val := d.Head.Val
	d.Head = d.Head.Next
    if d.Head == nil {
		d.Tail = nil
	} else {
		d.Head.Prev = nil
	}
	d.Count -= 1
	return true
}

// RemoveNode method
// O(N) time, O(1) space
func (d *DoublyLinkedList) RemoveNode(val interface{}) bool {
	curr := d.Head
	if curr == nil {
		return false
	}
	if curr.Val == val {
		curr = curr.Next
		d.count -= 1
		if curr != nil {
			d.Head = curr
			d.Head.Prev = nil
		} else {
			d.Tail = nil
		}
		return true
	}
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
			d.Count -= 1
			if curr.Next == nil {
				d.Tail = curr
			} else {
				d.Tail.Prev = curr
			}
			return true
		}
		curr = curr.Next
	}
	return false
}
