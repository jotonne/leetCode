package main

import (
	"fmt"
	"strings"
)

func main() {
	vowel := "aeiou"
	s := "aeiou"
	//    11111
	k := 2
	intList := []int{}
	for i := 0; i < len(s); i++ {
		if strings.Contains(vowel, string(s[i])) {
			intList = append(intList, 1)
		} else {
			intList = append(intList, 0)
		}
	}

	count := 0
	for i := 0; i < k; i++ {
		count += intList[i]
	}

	m := count
	for i := k; i < len(s); i++ {
		var first int
		if intList[i-k] == 0 {
			first = 0
		} else {
			first = -1
		}

		count = count + first + intList[i]
		m = max(m, count)
	}

	fmt.Println(m)
}
