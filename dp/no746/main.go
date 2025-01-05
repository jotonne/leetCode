package main

import "fmt"

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}

func minCostClimbingStairs(cost []int) int {
	var first, second int
	for _, v := range cost {
		first, second = second, min(first, second)+v
	}
	return min(first, second)
}
