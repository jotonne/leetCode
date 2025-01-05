package main

import "fmt"

func main() {
	n := 3
	fmt.Println(climbStairs(n))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	first := 1
	second := 2
	var temp int
	for i := 3; i <= n; i++ {
		temp = first + second
		first = second
		second = temp
	}
	return second
}
