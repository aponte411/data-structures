package non_linear

import (
	"fmt"
	"math"
	"testing"
)

func TestBinaryTree(t *testing.T) {
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
	fmt.Printf("Is tree valid binary search tree?: %v\n", IsBinarySearchTree(tree.Root, math.MinInt64, math.MaxInt64))
	fmt.Printf("Is tree balanced?: %v\n", IsBalanced(tree.Root))
	fmt.Printf("In order traversal: %v\n", InOrderTraversal(tree.Root))

}
