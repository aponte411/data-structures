package non_linear

import (
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := CreateBinarySearchTree(arr)
	tree.Add(66)
	tree.InOrderBST()
	fmt.Printf("Is 66 present in BST? %v\n", tree.FindRecursive(66))
	fmt.Printf("Min %v\n", tree.FindMinNode().Val)
	fmt.Printf("Max %v\n", tree.FindMaxNode().Val)
	fmt.Printf("Is balanced? %v\n", IsBSTBalanced(tree.Root))
	fmt.Printf("Second largest node %v\n", findSecondLargest(tree.Root))

}
