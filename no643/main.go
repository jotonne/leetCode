package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4

	var temp float64
	for i := 0; i < k; i++ {
		temp += float64(nums[i])
	}

	m := temp
	for i := k; i < len(nums); i++ {
		temp = temp - float64(nums[i-k]) + float64(nums[i])
		m = max(m, temp)
	}
	fmt.Println(m / float64(k))
}
