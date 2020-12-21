package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func CreateBinarySearchTree(arr []int) *BinarySearchTree {
	tree := new(BinarySearchTree)
	size := len(arr)
	tree.Root = createBinarySearchTree(arr, 0, size-1)
	return tree
}
func createBinarySearchTree(arr []int, start, end int) *Node {
	if start > end {
		return nil
	}
	mid := start + end // 2
	node := new(Node)
	node.Val = arr[mid]
	node.Left = createBinarySearchTree(arr, start, mid-1)
	node.Right = createBinarySearchTree(arr, mid+1, end)
	return node

}

func (b *BinarySearchTree) InOrderPrint() {
	fmt.Println("In-order")
	inOrderPrint(b.Root)
}
func inOrderPrint(node *Node) {
	if node == nil {
		return
	}
	inOrderPrint(node.Left)
	fmt.Println(node.Val)
	inOrderPrint(node.Right)
}

// TODO: keep in mind that new() returns a pointer to a value
// of type *T, which is equivalent to &T{}
// make(T) returns an initialized value of type T, which is
// for slices, maps, and channels

// O(N) time, O(N) space
func (b *BinarySearchTree) Add(val int) {
	b.Root = add(b.Root, val)
}
func add(node *Node, val int) *Node {
	if node == nil {
		node := &Node{val, nil, nil}
		return node
	}
	if val < node.Val {
		node.Left = add(node.Left, val)
	} else {
		node.Right = add(node.Right, val)
	}
	return node
}

// FindIterative
func (b *BinarySearchTree) FindIterative(val int) bool {
	node := b.Root
	for node != nil {
		if node.Val == val {
			return true
		} else if val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return false
}

// FindRecursive
// O(log n) average case time, O(H) worst case time
// O(log n) average case space, O(H) worst case space
func (b *BinarySearchTree) FindRecursive(val int) bool {
	return search(b.Root, val)
}
func search(node *Node, val int) bool {
	if node == nil {
		return true
	}
	if node.Val == val {
		return true
	} else if val < node.Val {
		return search(node.Left, val)
	} else {
		return search(node.Right, val)
	}
}

// FindMin
func (b *BinarySearchTree) FindMin() (int, bool) {
	node := b.Root
	if node == nil {
		fmt.Println("EmptyTreeException")
		return 0, false
	}
	for node.Left != nil {
		node = node.Left
	}
	return node.Val, true
}

// FindMinNode
func (b *BinarySearchTree) FindMinNode() *Node {
	node := b.Root
	if node == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// FindMax
func (b *BinarySearchTree) FindMax() (int, bool) {
	node := b.Root
	if node == nil {
		fmt.Println("EmptyTreeException")
		return 0, false
	}
	for node.Right != nil {
		node = node.Right
	}
	return node.Val, true
}

// SecondLargestNode
// Because we cant assume the BST is balanced, the parent of the rightmost parent
// isnt guaranteed to be the second largest, because the largest element can have a left
// subtree. This simplies the problem into two cases: the largest node has a left subtree or
// the largest node does not have a left subtree.
// So if the largest element does have a left subtree recurse through the left subtree
// and find the max. Else weve found the second largest node.
// O(h) time O(log n) if the BST is balanced; O(h) space (use iterative findmax for constant space)
func findSecondLargest(root *Node) int {
	curr := root
	for curr != nil {
		// If were at the largest node (rightmost) and there is a left subtree
		// recurse through left subtree and find second largest (max of left subtree)
		if curr.Left != nil && curr.Right == nil {
			return findMaxRecursive(curr.Left)
		}
		// If were at the second to last node and theres no left or right subtrees
		// we've found the second largest value
		if curr.Right != nil && curr.Right.Left == nil && curr.Right.Right == nil {
			return curr.Val
		}
		// Else the second largest is somewhere in that subtree
		curr = curr.Right
	}
	return curr.Val
}

// findMaxRecursive
// O(n) time, O(n) space, iterative has O(1) space since it traverses using
// variables, e.g. curr = curr.Right
func findMaxRecursive(root *Node) int {
	if root == nil {
		fmt.Println("EmptyTreeException")
		return 0
	}
	// The max node is the rightmost node so if there is a right
	// child, that node and the subtree below it are all greater than
	// the current, so step down to this child and recurse
	if root.Right != nil {
		return findMaxRecursive(root.Right)
	} else {
		return root.Val
	}
}

// FindMaxNode
func (b *BinarySearchTree) FindMaxNode() *Node {
	node := b.Root
	if node == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}

// height
func height(node *Node) int {
	if node == nil {
		return 0
	}
	left := height(node.Left)
	right := height(node.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

// IsBalanced
func IsBalanced(node *Node) bool {
	if node == nil {
		return true
	}
	left := height(node.Left)
	right := height(node.Right)
	return math.Abs(float64(left)-float64(right)) < 2 && IsBalanced(node.Left) && IsBalanced(node.Right)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := CreateBinarySearchTree(arr)
	tree.Add(66)
	tree.InOrderPrint()
	fmt.Printf("Is 66 present in BST? %v\n", tree.FindRecursive(66))
	fmt.Printf("Min %v\n", tree.FindMinNode().Val)
	fmt.Printf("Max %v\n", tree.FindMaxNode().Val)
	fmt.Printf("Is balanced? %v\n", IsBalanced(tree.Root))
	fmt.Printf("Second largest node %v\n", findSecondLargest(tree.Root))

}
