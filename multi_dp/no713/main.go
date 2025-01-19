package main

import "fmt"

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := -2
	fmt.Println(maxProfit(prices, fee))
}

func maxProfit(prices []int, fee int) int {
	holdDp := make([]int, len(prices))
	notHoldDp := make([]int, len(prices))

	holdDp[0] = -1 * prices[0]
	notHoldDp[0] = 0

	for i := 1; i < len(prices); i++ {
		holdDp[i] = max(holdDp[i-1], notHoldDp[i-1]-prices[i])
		notHoldDp[i] = max(notHoldDp[i-1], holdDp[i-1]+prices[i]-fee)
	}
	return max(holdDp[len(prices)-1], notHoldDp[len(prices)-1])
}
