package main

import (
	"fmt"
	"math"
)

func main() {
	num := 4
	fmt.Println(numTilings(num))
}

func numTilings(n int) int {
	dp := make([]int, 1000)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 5

	module := int(math.Pow(10, 9) + 7)
	for i := 4; i <= n; i++ {
		dp[i] = (2*dp[i-1] + dp[i-3]) % module
	}
	return dp[n]
}
