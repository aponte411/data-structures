package binary_search_tree

import (
	"fmt"
	"math"
)

type BinarySearchTreeNode struct {
	Val   int
	Left  *BinarySearchTreeNode
	Right *BinarySearchTreeNode
}

type BinarySearchTree struct {
	Root *BinarySearchTreeNode
}

func CreateBinarySearchTree(arr []int) *BinarySearchTree {
	tree := new(BinarySearchTree)
	size := len(arr)
	tree.Root = createBinarySearchTree(arr, 0, size-1)
	return tree
}

func createBinarySearchTree(arr []int, start, end int) *BinarySearchTreeNode {
	if start > end {
		return nil
	}
	mid := start + end // 2
	node := new(BinarySearchTreeNode)
	node.Val = arr[mid]
	node.Left = createBinarySearchTree(arr, start, mid-1)
	node.Right = createBinarySearchTree(arr, mid+1, end)
	return node

}

func (b *BinarySearchTree) InOrderBST() {
	fmt.Println("In-order")
	inOrderBST(b.Root)
}
func inOrderBST(node *BinarySearchTreeNode) {
	if node == nil {
		return
	}
	inOrderBST(node.Left)
	fmt.Println(node.Val)
	inOrderBST(node.Right)
}

// TODO: keep in mind that new() returns a pointer to a value
// of type *T, which is equivalent to &T{}
// make(T) returns an initialized value of type T, which is
// for slices, maps, and channels

// O(N) time, O(N) space
func (b *BinarySearchTree) Add(val int) {
	b.Root = add(b.Root, val)
}

func add(node *BinarySearchTreeNode, val int) *BinarySearchTreeNode {
	if node == nil {
		node := &BinarySearchTreeNode{val, nil, nil}
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
	return searchBST(b.Root, val)
}
func searchBST(node *BinarySearchTreeNode, val int) bool {
	if node == nil {
		return true
	}
	if node.Val == val {
		return true
	} else if val < node.Val {
		return searchBST(node.Left, val)
	} else {
		return searchBST(node.Right, val)
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
func (b *BinarySearchTree) FindMinNode() *BinarySearchTreeNode {
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
func findSecondLargest(root *BinarySearchTreeNode) int {
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
func findMaxRecursive(root *BinarySearchTreeNode) int {
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
func (b *BinarySearchTree) FindMaxNode() *BinarySearchTreeNode {
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
func height(node *BinarySearchTreeNode) int {
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
func IsBSTBalanced(node *BinarySearchTreeNode) bool {
	if node == nil {
		return true
	}
	left := height(node.Left)
	right := height(node.Right)
	return math.Abs(float64(left)-float64(right)) < 2 && IsBSTBalanced(node.Left) && IsBSTBalanced(node.Right)
}
