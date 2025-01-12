package main

import "fmt"

func main() {
	questions := [][]int{
		{3, 2},
		{4, 3},
		{4, 4},
		{2, 5},
	}
	fmt.Println(mostPoints(questions))
}

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n+1)

	for i := n - 1; n >= 0; i-- {
		val := int64(questions[i][0])
		brainpower := questions[i][1]

		if i+brainpower+1 > n {
			dp[i] = 0
		} else {
			dp[i] = max(dp[i+1], val+dp[i+brainpower+1])
		}
	}

	return dp[0]
}
