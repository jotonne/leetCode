package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}

	if len(nums) == 1 {
		return
	}

	var queue []int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			queue = append(queue, i)
			continue
		}
		if len(queue) == 0 {
			continue
		}
		index := queue[0]
		queue = queue[1:]
		nums[index] = nums[i]
		nums[i] = 0
		queue = append(queue, i)
	}
	fmt.Println(nums)
}
