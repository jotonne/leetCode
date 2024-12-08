package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	var resultsNums []int
	m1 := map[int]int{}
	m2 := map[int]int{}

	bak := 1
	for i := 0; i < len(nums); i++ {
		m1[i] = bak
		bak *= nums[i]
	}

	bak2 := 1
	for i := len(nums) - 1; i >= 0; i-- {
		m2[i] = bak2
		bak2 *= nums[i]
	}

	for i := 0; i < len(nums); i++ {
		resultsNums = append(resultsNums, m1[i]*m2[i])
	}

	fmt.Println(resultsNums)
}
