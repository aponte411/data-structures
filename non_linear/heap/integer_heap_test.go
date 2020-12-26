package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestIntegerHeap(t *testing.T) {

	intHeap := &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	res := (*intHeap)[0]
	if res != 1 {
		t.Errorf("Expected 1, got %v", res)
	}
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
