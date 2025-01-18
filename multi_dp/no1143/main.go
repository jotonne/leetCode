package main

import "fmt"

func main() {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence(text1, text2))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	results := make([][]int, m+1)
	for i := range results {
		results[i] = make([]int, n+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				results[i][j] = results[i-1][j-1] + 1
			} else {
				results[i][j] = max(results[i][j-1], results[i-1][j])
			}
		}
	}
	return results[m][n]
}
