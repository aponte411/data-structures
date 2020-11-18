package main

import (
	"container/heap"
	"fmt"
)

type IntegerHeap []int

// Length of integer heap
func (iheap IntegerHeap) Len() int {
	return len(iheap)
}

// Less
func (iheap IntegerHeap) Less(i, j int) bool {
	return iheap[i] < iheap[j]
}

// Swap two values
func (iheap IntegerHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i]
}

// Push
func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

// Pop
func (iheap *IntegerHeap) Pop() interface{} {
	prev := *iheap
	n := len(prev)
	// get last node
	node := prev[n-1]
	*iheap = prev[:n-1]
	return node
}

func main() {
	intHeap := &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("Min element: %v\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
