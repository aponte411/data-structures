package stack

import "fmt"

type StackNode struct {
	Val  interface{}
	Next *StackNode
}

type StackLinkedList struct {
	Head *StackNode
	size int
}

// IsEmpty
func (s *StackLinkedList) IsEmpty() bool { return s.size == 0 }

// Peek
func (s *StackLinkedList) Peek() interface{} {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return nil
	}
	return s.Head.Val
}

// Push
// O(1) time, O(1) space
func (s *StackLinkedList) Push(val interface{}) {
	node := &StackNode{val, s.Head}
	s.Head = node
	s.size += 1
}

// Pop
// O(1) time, O(1) space
func (s *StackLinkedList) Pop() interface{} {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return nil
	}
	val := s.Head.Val
	s.Head = s.Head.Next
	s.size -= 1
	return val
}

// Print
// O(N) time, O(1) space
func (s *StackLinkedList) Print() {
	for node := s.Head; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}

func IsBalancedParenthesis(expn string) bool {
	stk := new(StackLinkedList)
	for i := 0; i < len(expn); i++ {
		ch := string(expn[i])
		switch ch {
		case "{", "[", "(":
			stk.Push(ch)
		case "}":
			val := stk.Pop().(string)
			if val != "{" {
				return false
			}
		case "]":
			val := stk.Pop().(string)
			if val != "[" {
				return false
			}
		case ")":
			val := stk.Pop().(string)
			if val != "(" {
				return false
			}
		}
	}
	return stk.IsEmpty()
}
