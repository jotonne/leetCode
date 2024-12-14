package main

import "fmt"

func main() {
	s := "axc"
	t := "ahbgdc"

	index := 0
	count := 0
	for _, v := range s {
		for j := index; j < len(t); j++ {
			if string(v) == string(t[j]) {
				index = j + 1
				count++
				break
			}
		}
	}

	fmt.Println(count)
	// return count == len(s)

	// return true
}
