package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "  hello world  "
	strs := strings.Split(s, " ")

	var result []string
	for i := len(strs) - 1; i >= 0; i-- {
		if len(strs[i]) == 0 {
			continue
		}
		result = append(result, strs[i])
	}

	fmt.Println(strings.Join(result, " "))
}
