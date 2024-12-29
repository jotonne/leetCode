package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "leet**cod*e"
	var sl []string
	var r []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			if len(r) == 0 {
				continue
			}
			r = r[:len(r)-1]
		} else {
			r = append(r, s[i])
		}
	}
	// for _, v := range s {
	// 	if string(v) == "*" {
	// 		if len(sl) == 0 {
	// 			continue
	// 		}
	// 		sl = sl[:len(sl)-1]
	// 	} else {
	// 		sl = append(sl, string(v))
	// 	}
	// }
	fmt.Println(strings.Join(sl, ""))
	fmt.Println(string(r))
}
