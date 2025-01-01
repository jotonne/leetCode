package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	Node3 := TreeNode{}
	Node20 := TreeNode{}
	Node11 := TreeNode{
		Right: &Node20,
	}
	rootNode := TreeNode{
		Left:  &Node3,
		Right: &Node11,
	}

	fmt.Println(maxDepth(&rootNode))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDepth := maxDepth(root.Left)
	rDepth := maxDepth(root.Right)

	return 1 + max(lDepth, rDepth)
}
