package non_linear

import (
	"fmt"
	"testing"
)

func TestIntegerHeap(t *testing.T) {

	intHeap := &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("Min element: %v\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
