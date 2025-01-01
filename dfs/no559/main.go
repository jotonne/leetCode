package main

import (
	"fmt"
	"sort"
)

type Node struct {
	Val      int
	Children []*Node
}

func main() {
	Node3 := Node{}
	Node20 := Node{}
	rootNode := Node{
		Children: []*Node{&Node3, &Node20},
	}

	fmt.Println(maxDepth(&rootNode))
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	var depthList []int
	for _, n := range root.Children {
		depth := maxDepth(n)
		depthList = append(depthList, depth)
	}

	if len(depthList) == 0 {
		return 1
	} else {
		sort.Ints(depthList)
		return depthList[len(depthList)-1] + 1
	}
}

// other person answers
// func walk(c []*Node, depth int) int {
// 	if len(c) == 0 {
// 		return depth
// 	}
// 	maxDepth := depth
// 	for _, leaf := range c {
// 		d := walk(leaf.Children, depth+1)
// 		maxDepth = max(maxDepth, d)
// 	}
// 	return maxDepth
// }

// func maxDepth(root *Node) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return walk(root.Children, 1)
// }
