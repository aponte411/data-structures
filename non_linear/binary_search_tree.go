package main

import "fmt"

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

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := CreateBinarySearchTree(arr)
	tree.InOrderPrint()
}
