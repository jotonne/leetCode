package main

import "fmt"

func main() {
	m := 3
	n := 7
	fmt.Println(uniquePaths(m, n))
}

func uniquePaths(m int, n int) int {
	paths := make([][]int, m)
	for i := range paths {
		paths[i] = make([]int, n)
	}
	paths[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j-1 >= 0 {
				paths[i][j] += paths[i][j-1]
			}
			if i-1 >= 0 {
				paths[i][j] += paths[i-1][j]
			}
		}
	}

	fmt.Println(paths)
	return paths[m-1][n-1]
}
