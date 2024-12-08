package main

import (
	"fmt"
	"sort"
)

func main() {
	// var result []string{}
	test := []byte{97}

	var result []any

	sort.Slice(test, func(i, j int) bool { return test[i] < test[j] })

	var bak byte
	count := 0
	for i, v := range test {
		if bak == v {
			count++
		} else {
			if i != 0 {
				result = append(result, count)
			}
			result = append(result, v)
			bak = v
			count = 1
		}
	}
	result = append(result, count)
	fmt.Println(result)
	fmt.Println(len(result))
}
