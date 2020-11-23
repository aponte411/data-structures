package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type StackLinkedList struct {
	Head *Node
	size int
}

// IsEmpty
func (s *StackLinkedList) IsEmpty() bool { return s.size == 0 }

// Peek
func (s *StackLinkedList) Peek() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return 0, false
	}
	return s.Head.Val, true
}

// Push
// O(1) time, O(1) space
func (s *StackLinkedList) Push(val int) {
	node := &Node{val, s.Head}
	s.Head = node
	s.size += 1
}

// Pop
// O(1) time, O(1) space
func (s *StackLinkedList) Pop() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return 0, false
	}
	val := s.Head.Val
	s.Head = s.Head.Next
	s.size -= 1
	return val, true
}

// Print
// O(N) time, O(1) space
func (s *StackLinkedList) Print() {
	for node := s.Head; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}

func main() {
	s := &StackLinkedList{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for !s.IsEmpty() {
		val, _ := s.Pop()
		fmt.Println(val)
	}
	s.Push(10)
	s.Push(11)
	s.Push(12)
	s.Print()
}
