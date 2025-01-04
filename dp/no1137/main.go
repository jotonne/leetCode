package main

import "fmt"

func main() {
	n := 25
	fmt.Println(tribonacci(n))
}

func tribonacci(n int) int {
	var ns []int
	ns = append(ns, 0)
	ns = append(ns, 1)
	ns = append(ns, 1)

	if n < 3 {
		return ns[n]
	}

	for i := 3; i < n+1; i++ {
		temp := ns[i-3] + ns[i-2] + +ns[i-1]
		ns = append(ns, temp)
	}

	fmt.Println(ns)
	return ns[n]
}
