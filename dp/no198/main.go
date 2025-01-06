package main

import "fmt"

func main() {
	nums := []int{2, 1, 1, 2}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	first := nums[0]
	second := max(nums[0], nums[1])
	var temp int
	for i := 2; i < len(nums); i++ {
		temp = max(second, first+nums[i])
		first = second
		second = temp
	}
	return second
}
