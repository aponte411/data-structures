package main

import (
	"fmt"
)

type StackInt struct {
	items []int
}

func (s *StackInt) IsEmpty() bool {
	length := len(s.items)
	return length == 0
}

func (s *StackInt) Len() int {
	return len(s.items)
}

func (s *StackInt) Print() {
	for _, item := range s.items {
		fmt.Println(item)
	}
}

func (s *StackInt) Push(val int) {
	s.items = append(s.items, val)
}

func (s *StackInt) Pop() int {
	length := len(s.items)
	val := s.items[length-1]
	s.items = s.items[:length-1]
	return val
}

func (s *StackInt) Top() int {
	return s.items[len(s.items)-1]
}

func main() {
	s := &StackInt{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}
}
