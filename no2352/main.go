package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	grid := [][]int{
		{3, 2, 1},
		{1, 7, 6},
		{2, 7, 7},
	}
	// grid := [][]int{
	// 	{11, 1},
	// 	{1, 11},
	// }
	var h []string
	m := map[string]int{}
	for _, v := range grid {
		for _, n := range v {
			h = append(h, strconv.Itoa(n))
		}
		k := strings.Join(h, ",")
		m[k] = m[k] + 1
		h = nil
	}

	var h2 []string
	var result int
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid[0]); j++ {
			h2 = append(h2, strconv.Itoa(grid[j][i]))
		}
		k2 := strings.Join(h2, ",")
		result += m[k2]
		h2 = nil
	}
	fmt.Println(result)
}
