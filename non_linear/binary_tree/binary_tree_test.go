package binary_tree

import (
	"fmt"
	"math"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := LevelOrderBinaryTree(arr)
	tree.LevelOrderPrint()

	res1 := tree.SumAllNodes()
	if res1 != 55 {
		t.Errorf("Expected 55 got %v", res1)
	}

	res2 := tree.NumberOfLeaves()
	if res2 != 5 {
		t.Errorf("Expected 5 leaves, got %v", res2)
	}

	res3 := tree.Search(10)
	if !res3 {
		t.Errorf("10 should be in binary tree %v", res3)
	}

	res4 := tree.TreeDepth()
	if res4 != 4 {
		t.Errorf("Expected depth of 4, got %v", res4)
	}

	tree2 := tree.CopyBinaryTree()
	res5 := tree.IsEqual(tree2)
	if !res5 {
		t.Errorf("Trees should be equal, %v", res5)
	}

	fmt.Printf("Max Depth: %v\n", maxDepth(tree.Root))
	fmt.Printf("Is there a path sum equal to 7: %v\n", tree.HasPathSum(7))
	fmt.Printf("Is tree valid binary search tree?: %v\n", IsBinarySearchTree(tree.Root, math.MinInt64, math.MaxInt64))
	fmt.Printf("Is tree balanced?: %v\n", IsBTBalanced(tree.Root))
	fmt.Printf("In order traversal: %v\n", InOrderTraversal(tree.Root))

}
