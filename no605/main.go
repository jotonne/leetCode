package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "IceCreAm"
	var vowelsIndex []int
	for i, v := range s {
		if strings.Contains("aeiuo", strings.ToLower(string(v))) {
			vowelsIndex = append(vowelsIndex, i)
		}
	}
	runes := []rune(s)
	for i := 0; i < len(vowelsIndex)/2; i++ {
		runes[vowelsIndex[i]], runes[vowelsIndex[len(vowelsIndex)-1-i]] = runes[vowelsIndex[len(vowelsIndex)-1-i]], runes[vowelsIndex[i]]
	}

	fmt.Println(string(runes))
}
