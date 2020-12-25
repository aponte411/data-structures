package non_linear

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
