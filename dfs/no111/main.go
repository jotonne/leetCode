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

	fmt.Println(minDepth(&rootNode))
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lDepth := minDepth(root.Left)
	rDepth := minDepth(root.Right)

	if lDepth == 0 {
		return rDepth + 1
	} else if rDepth == 0 {
		return lDepth + 1
	} else {
		return min(lDepth, rDepth) + 1
	}
}
