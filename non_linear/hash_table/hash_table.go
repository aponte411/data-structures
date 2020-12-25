package non_linear

import "fmt"

// HashTable using separate chaining to deal with hash collisions
type HashTable struct {
	array []*Node
	size  int
}

type Node struct {
	key   int
	value int
	next  *Node
}

// Init method
func (h *HashTable) Init() {
	h.size = 769 // something prime
	h.array = make([]*Node, h.size)
	// zero value of nil instead of 0
	for i := 0; i < h.size; i++ {
		h.array[i] = nil
	}
}

// Hash function
func (h *HashTable) Hash(key int) int {
	return key % h.size
}

// Add
// O(1) time, O(1) space
func (h *HashTable) Add(key int, value int) {
	// hash key
	index := h.Hash(key)
	// instatiate new node with next pointer to itself
	node := &Node{key, value, h.array[index]}
	// assign value to array index
	h.array[index] = node
}

// Remove method
// O(1) time on avg but O(N) in worst case, O(1) time
func (h *HashTable) Remove(key int) bool {
	index := h.Hash(key)
	// get head of linked list
	head := h.array[index]
	// if key is head
	if head != nil && head.key == key {
		h.array[index] = head.next
		return true
	}
	// if key is down the list
	for head = head; head != nil; head = head.next {
		if head != nil && head.key == key {
			h.array[index] = head.next
			return true
		}
	}
	return false
}

// Find method
// O(1) time on avg, O(N) in the worst case
func (h *HashTable) Find(key int) bool {
	index := h.Hash(key)
	for head := h.array[index]; head != nil; head = head.next {
		if head.key == key {
			return true
		}
	}
	return false
}

// Get
func (h *HashTable) Get(key int) int {
	index := h.Hash(key)
	for head := h.array[index]; head != nil; head = head.next {
		if head.key == key {
			return head.value
		}
	}
	return 0
}

// Print
func (h *HashTable) Print() {
	for i := 0; i < h.size; i++ {
		head := h.array[i]
		if head != nil {
			fmt.Printf("Values at index %v\n", i)
			for head != nil {
				fmt.Printf("%v\n", head.value)
				head = head.next
			}
		}
	}
	fmt.Println()
}
