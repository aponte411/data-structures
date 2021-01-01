package linked_list

import (
	"fmt"
)

// Singly linked list node
type Node struct {
	Val  interface{}
	Next *Node
}

// Singly linked list
type SinglyLinkedList struct {
	Head  *Node
	Count int
}

// AddToHead method
func (s *SinglyLinkedList) AddToHead(val interface{}) {
	node := &Node{val, nil}
	s.Count += 1
	if s.Head != nil {
		node.Next = s.Head
	}
	s.Head = node
}

// RemoveHead method
// O(1) time, O(1) space
func (s *SinglyLinkedList) RemoveHead() bool {
	if s.Count == 0 {
		return false
	}
	s.Head = s.Head.Next
	s.Count -= 1
	return true

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
func (s *SinglyLinkedList) AddToEnd(val interface{}) {
	node := &Node{val, nil}
	s.Count += 1
	last := s.LastNode()
	if last != nil {
		last.Next = node
	}
}

// Find method
// O(N) time, O(1) space
func (s *SinglyLinkedList) Find(val interface{}) bool {
	for node := s.Head; node != nil; node = node.Next {
		if node.Val == val {
			return true
		}
	}
	return false
}

// NodeWithValue method
// O(N) time, O(1) space
func (s *SinglyLinkedList) NodeWithValue(val interface{}) *Node {
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
func (s *SinglyLinkedList) AddAfter(nodeVal, val interface{}) {
	node := &Node{val, nil}
	s.Count += 1
	nodeWith := s.NodeWithValue(nodeVal)
	if nodeWith != nil {
		node.Next = nodeWith.Next
		nodeWith.Next = node
	}
}

// DeleteNode method
// O(N) time and O(1) space
func (s *SinglyLinkedList) DeleteNode(val interface{}) bool {
	// if list is empty return false
	if s.Count == 0 {
		return false
	}
	if val == s.Head.Val {
		s.Head = s.Head.Next
		s.Count -= 1
		return true
	}

	for node := s.Head; node != nil; node = node.Next {
		if node.Next.Val == val {
			node.Next = node.Next.Next
			s.Count -= 1
			return true
		}
	}
	return false
}

// Reverse Print
func (s *SinglyLinkedList) ReversePrint(node *Node) {
	if node == nil {
		return
	}
	s.ReversePrint(node.Next)
	fmt.Println(node.Val)
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

// RemoveDuplicates
// O(N) time, O(1) space
func (s *SinglyLinkedList) RemoveDuplicates() {
	for curr := s.Head; curr != nil; curr = curr.Next {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
			s.Count -= 1
			return
		}
	}
}

// HasCycle method
func (s *SinglyLinkedList) HasCycle() bool {
	fast, slow := s.Head, s.Head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// KtoLastNode method
// O(N-k) time, O(1) space
func (s *SinglyLinkedList) KtoLastNode(k int) *Node {
	node := s.Head
	kToLast := s.Count - k
	for i := 0; i < kToLast; i++ {
		node = node.Next
	}
	return node
}

// RemoveNthFromEnd is an alternative method to KtoLastNode
// O(N) time, O(1) space
func (s *SinglyLinkedList) RemoveNthFromEnd(n int) *Node {
	dummy := &Node{0, nil}
	dummy = s.Head
	first := dummy
	second := dummy
	for i := 1; i <= n+1; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
