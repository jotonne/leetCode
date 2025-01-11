package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxNum := nums[0]
	minNum := nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		var temp []int
		temp = append(temp, nums[i])
		temp = append(temp, maxNum*nums[i])
		temp = append(temp, minNum*nums[i])
		sort.Ints(temp)
		fmt.Println(temp)
		minNum = temp[0]
		maxNum = temp[len(temp)-1]
		result = max(result, maxNum)
	}
	return result
}
