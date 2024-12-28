package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 1, 1, 3}
	m := map[int]int{}
	for _, v := range arr {
		_, ok := m[v]
		if ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}

	r := true
	m2 := map[int]struct{}{}
	for _, v := range m {
		if _, ok := m2[v]; !ok {
			m2[v] = struct{}{}
			continue
		} else {
			r = false
			break
		}
	}

	fmt.Println(r)
}
