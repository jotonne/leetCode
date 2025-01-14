package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	results := make([][]int, m)
	for i := range results {
		results[i] = make([]int, n)
	}

	results[0][0] = grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if j-1 >= 0 && i-1 >= 0 {
				results[i][j] = min(results[i][j-1], results[i-1][j]) + grid[i][j]
				continue
			}
			if j-1 >= 0 && i-1 < 0 {
				results[i][j] = results[i][j-1] + grid[i][j]
				continue
			}
			if j-1 < 0 && i-1 >= 0 {
				results[i][j] = results[i-1][j] + grid[i][j]
				continue
			}
		}
	}
	return results[m-1][n-1]
}
