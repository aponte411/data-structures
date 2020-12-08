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

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := CreateBinarySearchTree(arr)
	tree.Add(66)
	tree.InOrderPrint()
	fmt.Printf("Is 66 present in BST? %v\n", tree.FindRecursive(66))
}