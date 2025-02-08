package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	Node3 := TreeNode{}
	// Node20 := TreeNode{}
	// Node11 := TreeNode{
	// 	Right: &Node20,
	// }
	rootNode := TreeNode{
		Left: &Node3,
	}

	fmt.Println(longestZigZag(&rootNode))
}

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(dfs(root.Left, true, 0), dfs(root.Right, false, 0))
}

func dfs(root *TreeNode, beforeLeft bool, zigzagCount int) int {
	if root == nil {
		return zigzagCount
	}

	var leftVal, rightVal int
	if beforeLeft {
		leftVal = max(dfs(root.Right, false, zigzagCount+1), dfs(root.Left, true, 0))
	} else {
		rightVal = max(dfs(root.Left, true, zigzagCount+1), dfs(root.Right, false, 0))
	}

	return max(leftVal, rightVal)
}
