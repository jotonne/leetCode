package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	Node1 := TreeNode{
		Val: 1,
	}
	Node3 := TreeNode{
		Val: 3,
	}
	Node2 := TreeNode{
		Left:  &Node1,
		Right: &Node3,
		Val:   2,
	}
	Node7 := TreeNode{
		Val: 7,
	}
	rootNode := TreeNode{
		Left:  &Node2,
		Right: &Node7,
		Val:   4,
	}

	fmt.Println(searchBST(&rootNode, 7))
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val == root.Val {
		return root
	}

	if val > root.Val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}

// 別解
// func searchBST(root *TreeNode, val int) *TreeNode {
// 	cp := root
// 	for cp.Val != val {
// 		if cp == nil {
// 			cp = nil
// 		}
// 		if val > cp.Val {
// 			cp = cp.Right
// 		} else {
// 			cp = cp.Left
// 		}
// 	}
// 	return cp
// }
