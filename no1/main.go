package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6

	m := map[int]int{}

	var ans []int
	for i, v := range nums {
		temp := target - v
		index, ok := m[v]
		if ok {
			ans = append(ans, index)
			ans = append(ans, i)
		} else {
			m[temp] = i
		}
	}

	fmt.Println(m)
	fmt.Println(ans)
}
