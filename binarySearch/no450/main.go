package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node2 := &TreeNode{Val: 2, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 4, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 3, Left: node2, Right: node4}

	node7 := &TreeNode{Val: 7, Left: nil, Right: nil}
	node6 := &TreeNode{Val: 6, Left: nil, Right: node7}

	rootNode := &TreeNode{Val: 5, Left: node3, Right: node6}
	// rootNode0 := &TreeNode{Val: 0}

	fmt.Println(deleteNode(rootNode, 5))
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key == root.Val {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left != nil && root.Right != nil {
			baseNode := root.Left
			for baseNode.Right != nil {
				baseNode = baseNode.Right
			}
			maxValue := baseNode.Val
			root.Left = deleteNode(root.Left, maxValue)
			root.Val = maxValue
			return root
		}
		if root.Left != nil {
			return root.Left
		}
		if root.Right != nil {
			return root.Right
		}
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}
