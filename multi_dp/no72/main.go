package main

import "fmt"

func main() {
	word1 := "input"
	word2 := "import"
	fmt.Println(minDistance(word1, word2))
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 0; i <= len(word1); i++ {
		dp[i][0] = i
	}

	for i := 0; i <= len(word2); i++ {
		dp[0][i] = i
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			up := dp[i-1][j] + 1
			left := dp[i][j-1] + 1

			var leftUp int
			if word1[i-1] == word2[j-1] {
				leftUp = dp[i-1][j-1]
			} else {
				leftUp = dp[i-1][j-1] + 1
			}

			dp[i][j] = min(up, min(left, leftUp))
		}
	}

	return dp[len(word1)][len(word2)]
}
