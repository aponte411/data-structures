package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

// Create a complete binary tree
// O(N) time, O(N) space
func LevelOrderBinaryTree(arr []int) *BinaryTree {
	tree := new(BinaryTree)
	tree.Root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func levelOrderBinaryTree(arr []int, start, size int) *Node {
	curr := &Node{arr[start], nil, nil}

	left := 2*start + 1
	right := 2*start + 2

	if left < size {
		curr.Left = levelOrderBinaryTree(arr, left, size)
	}
	if right < size {
		curr.Right = levelOrderBinaryTree(arr, right, size)
	}
	return curr
}

// Pre-order
func (b *BinaryTree) PreOrderPrint() {
	fmt.Println("Pre-order")
	preOrderPrint(b.Root)
}

func preOrderPrint(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preOrderPrint(root.Left)
	preOrderPrint(root.Right)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := LevelOrderBinaryTree(arr)
	tree.PreOrderPrint()
}
