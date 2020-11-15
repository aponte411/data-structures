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
	node := &Node{val, nil}
	s.count += 1
	if s.Head != nil {
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
// O(N) time, O(1) space
func (s *SinglyLinkedList) Iterate() {
	for node := s.Head; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
	fmt.Println("")
}

// LastNode method
// O(N) time, O(1) space
func (s *SinglyLinkedList) LastNode() *Node {
	var last *Node
	for node := s.Head; node != nil; node = node.Next {
		if node == nil {
			last = node
		}
	}
	return last
}

// AddToEnd method
// O(N) time, O(1) space
func (s *SinglyLinkedList) AddToEnd(val int) {
	node := &Node{val, nil}
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
	for node := s.Head; node != nil; node = node.Next {
		if node.Val == val {
			return true
		}
	}
	return false
}

// NodeWithValue method
// O(N) time, O(1) space
func (s *SinglyLinkedList) NodeWithValue(val int) *Node {
	var nodeWith *Node
	for node := s.Head; node != nil; node = node.Next {
		if node.Val == val {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddAfter method
func (s *SinglyLinkedList) AddAfter(nodeVal, val int) {
	node := &Node{val, nil}
	s.count += 1

	var nodeWith *Node
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
		if node.Next.Val == val {
			node.Next = node.Next.Next
			s.count -= 1
			return true
		}
	}
	return false
}

// Reverse method
// O(N) time, O(1) space
func (s *SinglyLinkedList) Reverse() {
	curr := s.Head
	var prev, next *Node
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	s.Head = prev
}

// Recursively reverse
// O(N) time, O(1) space
func (s *SinglyLinkedList) ReverseRecursively() {
	s.Head = s.reverseHelper(s.Head, nil)
}

func (s *SinglyLinkedList) reverseHelper(curr *Node, next *Node) *Node {
	var resultNode *Node
	if curr == nil {
		return nil
	}
	if curr.Next == nil {
		curr.Next = next
		return curr
	}
	resultNode = s.reverseHelper(curr.Next, curr)
	curr.Next = next
	return resultNode
}


// SortedInsert
// O(N) time, O(1) space
func (s *SinglyLinkedList) SortedInsert(val int) {
	node := &Node{val, nil}
	curr := s.Head
	if curr == nil || curr.Val > val {
		node.Next = s.Head
		s.Head = node
		return
	}
	for curr.Next != nil && curr.Next.Val < val {
		curr = curr.Next
	}
	node.Next = curr.Next
	curr.Next = node
	s.count += 1
}

// RemoveDuplicates
// O(N) time, O(1) space
func (s *SinglyLinkedList) RemoveDuplicates() {
	for curr := s.Head; curr != nil; curr = curr.Next {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
			s.count -= 1
			return
		}
	}
}

// main
func main() {
	list := &SinglyLinkedList{}
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
    list.Reverse()
}
