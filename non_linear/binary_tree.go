package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
	"math"
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
	// or node := new(Node)
	// node.Val = arr[start]
	// node.Left = nil
	// node.Right = nil
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

// Iterative InOrderTraversal
func InOrderTraversal(root *Node) []int {
	if root == nil {
		return nil
	}
	stack := make([]*Node, 0)
	res := make([]int, 0)
	curr := root
	for curr != nil || len(stack) != 0 {
		// go all the way to the left
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		// pop the stack
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// add to results
		res = append(res, curr.Val)
		// go to the right
		curr = curr.Right
	}
	return res
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

// SumAllNode
func (b *BinaryTree) SumAllNodes() int {
	return sumAllNodes(b.Root)
}

func sumAllNodes(node *Node) int {
	if node == nil {
		return 0
	}
	return sumAllNodes(node.Left) + sumAllNodes(node.Right) + node.Val
}

// NumberOfLeaves
func (b *BinaryTree) NumberOfLeaves() int {
	return numberOfLeaves(b.Root)
}
func numberOfLeaves(node *Node) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}
	return numberOfLeaves(node.Left) + numberOfLeaves(node.Right)
}

// Search
func (b *BinaryTree) Search(val int) bool {
	return search(b.Root, val)
}

func search(node *Node, val int) bool {
	var left, right bool
	if node == nil {
		return false
	}
	if node.Val == val {
		return true
	}
	left = search(node.Left, val)
	if left {
		return true
	}
	right = search(node.Right, val)
	if right {
		return true
	}
	return false
}

// TreeDepth
func (b *BinaryTree) TreeDepth() int {
	return treeDepth(b.Root)
}

func treeDepth(node *Node) int {
	if node == nil {
		return 0
	}
	left := treeDepth(node.Left)
	right := treeDepth(node.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}

}

// IsBalancedTree
func IsBalanced(node *Node) bool {
	if node == nil {
		return true
	}
	left := treeDepth(node.Left)
	right := treeDepth(node.Right)
	return math.Abs(float64(left)-float64(right)) < 2 && IsBalanced(node.Left) && IsBalanced(node.Right)
}

// TODO: NEED TO FIX
// FindMax
// O(N) time, O(1)
func (b *BinaryTree) FindMax() int {
	return findMax(b.Root)
}
func findMax(node *Node) int {
	if node == nil {
		return math.MinInt32
	}
	max := node.Val
	left := treeDepth(node.Left)
	if left > max {
		max = left
	}
	right := treeDepth(node.Right)
	if right > max {
		max = right
	}
	return max
}

// IsEqual
func (b *BinaryTree) IsEqual(t2 *BinaryTree) bool {
	return isEqual(b.Root, t2.Root)
}
func isEqual(node1 *Node, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.Val == node2.Val && isEqual(node1.Left, node2.Left) && isEqual(node1.Right, node2.Right)
}

// CopyTree
func (b *BinaryTree) CopyTree() *BinaryTree {
	tree2 := new(BinaryTree)
	tree2.Root = copyTree(b.Root)
	return tree2
}
func copyTree(node *Node) *Node {
	var newNode *Node
	if node != nil {
		newNode = new(Node)
		newNode.Val = node.Val
		newNode.Left = copyTree(node.Left)
		newNode.Right = copyTree(node.Right)
		return newNode
	}
	return nil
}

// HasPathSum
func (b *BinaryTree) HasPathSum(sum int) bool {
	hasPath := hasPathSum(b.Root, sum)
	if hasPath {
		return true
	} else {
		return false
	}
}
func hasPathSum(node *Node, sum int) bool {
	if node == nil {
		return false
	}
	sum -= node.Val
	if node.Left == nil && node.Right == nil {
		return sum == 0
	}
	left := hasPathSum(node.Left, sum)
	right := hasPathSum(node.Right, sum)
	return left || right
}

// MaxDepth
func maxDepth(node *Node) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}
	left := maxDepth(node.Left)
	right := maxDepth(node.Right)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

// IsBinarySearchTree
func IsBinarySearchTree(node *Node, min int, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}
	return IsBinarySearchTree(node.Left, min, node.Val) && IsBinarySearchTree(node.Right, node.Val, max)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := LevelOrderBinaryTree(arr)
	tree.LevelOrderPrint()
	fmt.Printf("Sum of all nodes: %d\n", tree.SumAllNodes())
	fmt.Printf("Number of leaves: %d\n", tree.NumberOfLeaves())
	fmt.Printf("Is 10 in binary tree? %v\n", tree.Search(10))
	fmt.Printf("Binary tree depth: %v\n", tree.TreeDepth())
	tree2 := tree.CopyTree()
	fmt.Printf("Trees Equal?: %v\n", tree.IsEqual(tree2))
	fmt.Printf("Max Depth: %v\n", maxDepth(tree.Root))
	fmt.Printf("Is there a path sum equal to 7: %v\n", tree.HasPathSum(7))
	fmt.Printf("Is tree valid binary search tree?: %v\n", IsBinarySearchTree(tree.Root, MinInt, MaxInt))
	fmt.Printf("Is tree balanced?: %v\n", IsBalanced(tree.Root))
	fmt.Printf("In order traversal: %v\n", InOrderTraversal(tree.Root))

}
