package main

import (
	"fmt"
)

// Singly linked list node
type Node struct {
	Val  int
	Next *Node
}

// Singly linked list
type SinglyLinkedList struct {
	Head  *Node
	count int
}

// AddToHead method
func (s *SinglyLinkedList) AddToHead(val int) {
	var node = &Node{}
	node.Val = val
	node.Next = nil
	s.count += 1

	if node.Head != nil {
		node.Next = s.Head
	}
	s.Head = node
}

// RemoveHead method
// O(1) time, O(1) space
func (s *SinglyLinkedList) RemoveHead() (int, bool) {
	if s.count == 0 {
		return 0, false
	}
	val := s.Head.Val
	s.Head = s.Head.Next
	s.count -= 1
	return val, true

}

// Iterate through nodes and print to stdout
func (s *SinglyLinkedList) Iterate() {
	var node = *Node
	for node = s.Head; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}

// LastNode method
func (s *SinglyLinkedList) LastNode() *Node {
	var node *Node
	var last *Node
	for node = s.Head; node != nil; node = node.Next {
		if node == nil {
			last = node
		}
	}
	return last
}

// AddToEnd method
func (s *SinglyLinkedList) AddToEnd(val int) {
	var node = &Node{}
	node.Val = val
	node.Next = nil
	s.count += 1

	var last *Node
	last = s.LastNode()

	if last != nil {
		last.Next = node
	}
}

// Find method
// O(N) time, O(1) space
func (s *SinglyLinkedList) Find(val int) bool {
	node := s.Head
	for node != nil {
		if node.Val == val {
			return true
		}
		node = node.Next
	}
	return false
}

// NodeWithValue method
func (s *SinglyLinkedList) NodeWithValue(val int) *Node {
	var nodeWith = *Node
	var node = *Node
	for node = s.Head; node != nil; node = node.Next {
		if node.Val == val {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddAfter method
func (s *SinglyLinkedList) AddAfter(nodeVal, val int) {
	var node = &Node
	node.Val = val
	node.Next = nil
	s.count += 1

	var nodeWith = *Node
	nodeWith = s.NodeWithValue(nodeVal)
	if nodeWith != nil {
		node.Next = nodeWith.Next
		nodeWith.Next = node
	}
}

// DeleteNode method
// O(N) time and O(1) space
func (s *SinglyLinkedList) DeleteNode(val int) bool {
	// if list is empty return false
	if s.count == 0 {
		return false
	}
	if val == s.Head.Val {
		s.Head = s.Head.Next
		s.count -= 1
		return true
	}

	for node := s.Head; node != nil; node = node.Next {
		if node.next.Val == val {
			node.Next = node.Next.Next
			s.count -= 1
			return true
		}
	}
	return false
}

// main
func main() {
	var list SinglyLinkedList
	list = SinglyLinkedList{}
	list.AddToHead(1)
	list.AddToHead(2)
	list.AddToEnd(4)
	list.AddAfter(1, 9)
	list.Iterate()
}
