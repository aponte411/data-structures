package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
)

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

// Post-order
func (b *BinaryTree) PostOrderPrint() {
	fmt.Println("Post-order")
	postOrderPrint(b.Root)
}

func postOrderPrint(root *Node) {
	if root == nil {
		return
	}
	postOrderPrint(root.Left)
	postOrderPrint(root.Right)
	fmt.Println(root.Val)
}

// In-order
func (b *BinaryTree) InOrderPrint() {
	fmt.Println("In-order")
	inOrderPrint(b.Root)
}

func inOrderPrint(root *Node) {
	if root == nil {
		return
	}
	inOrderPrint(root.Left)
	fmt.Println(root.Val)
	inOrderPrint(root.Right)
}

// BFS with queue
func (b *BinaryTree) BFSPrint() {
	queue := new(queue.Queue)
	var node *Node
	if b.Root != nil {
		queue.Enqueue(b.Root)
	}
	for queue.Len() != 0 {
		tmp := queue.Dequeue()
		node = tmp.(*Node)
		fmt.Println(node.Val)
		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}
}

// DFS with stack
func (b *BinaryTree) DFSPrint() {
	stack := new(stack.Stack)
	if b.Root != nil {
		stack.Push(b.Root)
	}
	for stack.Len() != 0 {
		node := stack.Pop().(*Node)
		fmt.Println(node.Val)
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
}

// LevelOrderPrint

func (b *BinaryTree) LevelOrderPrint() {
	queue := new(queue.Queue)
	if b.Root != nil {
		queue.Enqueue(b.Root)
	}

	for queue.Len() != 0 {
		n := queue.Len()
		for i := 0; i < n; i++ {
			node := queue.Dequeue().(*Node)
			fmt.Println(node.Val)
			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := LevelOrderBinaryTree(arr)
	tree.LevelOrderPrint()
}
