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

// AddToHead method
// O(1) time, O(1) space
func (d *DoublyLinkedList) AddToHead(val interface{}) {
	node := &DoublyLinkedNode{val, nil, nil}
	if d.Count == 0 {
		d.Tail = node
		d.Head = node
	} else {
		d.Head.Prev = node
		node.Next = d.Head
		d.Head = node
	}
	d.Count += 1
}

func (d *DoublyLinkedList) AddToTail(val interface{}) {
	node := &DoublyLinkedNode{val, nil, nil}
	if d.Count == 0 {
		d.Tail = node
		d.Head = node
	} else {
		d.Tail.Next = node
		node.Prev = d.Tail
		d.Tail = node
	}
	d.Count += 1
}

// RemoveHead
// O(1) time, O(1) space
func (d *DoublyLinkedList) RemoveHead() bool {
	if d.Count == 0 {
		return false
	}
	node := d.Head
	if node.Next != nil {
		d.Head = d.Head.Next
	}

	d.Count -= 1
	return true
}

func (d *DoublyLinkedList) RemoveTail() bool {
	if d.Count == 0 {
		return false
	}
	node := d.Tail
	if node.Prev != nil {
		d.Tail = node.Prev
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
		d.Count -= 1
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
