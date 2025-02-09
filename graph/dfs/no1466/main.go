package main

import "fmt"

func main() {
	n := 6
	connections := [][]int{
		{0, 1},
		{1, 3},
		{2, 3},
		{4, 0},
		{4, 5},
	}
	fmt.Println(minReorder(n, connections))
}

func minReorder(n int, connections [][]int) int {
	adList := make(map[int][]int)

	for _, connection := range connections {
		adList[connection[0]] = append(adList[connection[0]], connection[1])
		adList[connection[1]] = append(adList[connection[1]], -connection[0])
	}

	hasArrived := make(map[int]bool)

	var count int
	var dfs func(city int)
	dfs = func(city int) {
		if hasArrived[city] {
			return
		} else {
			hasArrived[city] = true
		}

		for _, arriveCity := range adList[city] {
			if arriveCity > 0 {
				if !hasArrived[arriveCity] {
					count++
				}
				dfs(arriveCity)
			} else {
				dfs(-arriveCity)
			}
		}
	}

	dfs(0)
	return count
}
