package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	Node3 := TreeNode{
		Val: 3,
	}
	Node20 := TreeNode{
		Val: 20,
	}
	Node11 := TreeNode{
		Val:   11,
		Right: &Node20,
	}
	rootNode := TreeNode{
		Val:   1,
		Left:  &Node3,
		Right: &Node11,
	}

	lLeafs := leafList(&rootNode)
	rLeafs := leafList(&rootNode)
	fmt.Println(reflect.DeepEqual(lLeafs, rLeafs))
}

func leafList(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	lLeafs := leafList(root.Left)
	rLeafs := leafList(root.Right)

	if lLeafs == nil && rLeafs == nil {
		return []int{root.Val}
	}
	return append(lLeafs, rLeafs...)
}
