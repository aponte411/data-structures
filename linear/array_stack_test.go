package linear

import "testing"

func TestStack(t *testing.T){
	s := new(StackInt)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}
    if !s.IsEmpty() {
        t.Errorf("Expected empty stack")
    }

}
