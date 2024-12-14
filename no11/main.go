package main

import (
	"fmt"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	m := -1
	left := 0
	right := len(height) - 1

	for width := len(height) - 1; width > 0; width-- {
		if height[left] > height[right] {
			m = max(m, width*height[right])
			right--
		} else {
			m = max(m, width*height[left])
			left++
		}
	}

	fmt.Println(m)
}
