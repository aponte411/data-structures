package binary_search_tree

import (
	"testing"
)

func TestBST(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := CreateBinarySearchTree(arr)
	tree.Add(66)
	tree.InOrderBST()

	res1 := tree.FindRecursive(66)
	if !res1 {
		t.Errorf("66 should be present %v", res1)
	}

	res2 := tree.FindMinNode().Val
	if res2 != 1 {
		t.Errorf("Expected 1, got %v", res2)
	}

	res3 := tree.FindMaxNode().Val
	if res3 != 66 {
		t.Errorf("Expected 66, got %v", res3)
	}
	res4 := IsBSTBalanced(tree.Root)

	if res4 {
		t.Errorf("BST should be unbalanced %v", res4)
	}

	res5 := findSecondLargest(tree.Root)
	if res5 != 10 {
		t.Errorf("Expected 10, got %v", res5)
	}

}
