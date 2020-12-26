package linear

import (
	"fmt"
	"testing"
)

func TestStackLinkedList(t *testing.T) {

	s := new(StackLinkedList)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for !s.IsEmpty() {
		val := s.Pop()
		fmt.Println(val)
	}
	s.Push(10)
	s.Push(11)
	s.Push(12)
	res := s.Pop()
	if res != 12 {
		t.Errorf("Expected 12, got %v", res)
	}
}
