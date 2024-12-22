package main

import (
	"fmt"
)

func main() {
	gain := []int{-5, 1, 5, 0, -7}

	sums := make([]int, len(gain)+1)
	sums[0] = 0
	for i, v := range gain {
		sums[i+1] = sums[i] + v
	}

	max := sums[0]
	for _, v := range sums {
		if max < v {
			max = v
		}
	}

	fmt.Println(max)
}
