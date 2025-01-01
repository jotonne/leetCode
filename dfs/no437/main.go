package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 1}

	// root := &TreeNode{Val: 5}
	// root.Left = &TreeNode{Val: 4}
	// root.Right = &TreeNode{Val: 8}

	// root.Left.Left = &TreeNode{Val: 11}
	// root.Right.Left = &TreeNode{Val: 13}
	// root.Right.Right = &TreeNode{Val: 4}

	// root.Left.Left.Left = &TreeNode{Val: 7}
	// root.Left.Left.Right = &TreeNode{Val: 2}

	// root.Right.Right.Left = &TreeNode{Val: 5}
	// root.Right.Right.Right = &TreeNode{Val: 1}

	pastNums := []int{}
	targetSum := 1
	count := calculate(root, targetSum, pastNums)
	fmt.Println(count)
}

func calculate(root *TreeNode, targetSum int, pastNums []int) int {
	if root == nil {
		return 0
	}

	var count int
	sum := root.Val
	if sum == targetSum {
		count++
	}
	for i := len(pastNums) - 1; i >= 0; i-- {
		sum += pastNums[i]
		if sum == targetSum {
			count++
		}
	}

	pastNums = append(pastNums, root.Val)
	count += calculate(root.Left, targetSum, pastNums)
	count += calculate(root.Right, targetSum, pastNums)
	return count
}
