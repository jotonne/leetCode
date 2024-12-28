package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}

	m1 := map[int]struct{}{}
	m2 := map[int]struct{}{}
	for _, v := range nums1 {
		m1[v] = struct{}{}
	}
	for _, v := range nums2 {
		m2[v] = struct{}{}
	}

	m12 := map[int]struct{}{}
	m22 := map[int]struct{}{}
	var a1, a2 []int
	for _, v := range nums1 {
		if _, ok := m2[v]; !ok {
			if _, ok := m12[v]; !ok {
				a1 = append(a1, v)
				m12[v] = struct{}{}
			}
		}
	}
	for _, v := range nums2 {
		if _, ok := m1[v]; !ok {
			if _, ok := m22[v]; !ok {
				a2 = append(a2, v)
				m22[v] = struct{}{}
			}
		}
	}

	var a3 [][]int
	a3 = append(a3, a1, a2)

	fmt.Println(a3)
}
