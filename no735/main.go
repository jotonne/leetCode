package main

import (
	"fmt"
	"math"
)

func main() {
	asteroids := []int{10, 2, -5}
	var r []int
	if len(asteroids) == 1 {
		// return []int{asteroids[0]}
	}
	for i := 0; i < len(asteroids); i++ {
		if len(r) == 0 {
			r = append(r, asteroids[i])
			continue
		}

		for j := len(r) - 1; j >= 0; j-- {
			latest := r[j]
			if asteroids[i] < 0 && latest > 0 {
				if math.Abs(float64(asteroids[i])) == math.Abs(float64(latest)) {
					r = r[:j]
					break
				} else if math.Abs(float64(asteroids[i])) > math.Abs(float64(latest)) {
					r = r[:j]
					if j == 0 {
						r = append(r, asteroids[i])
					}
					continue
				} else {
					break
				}
			} else {
				r = append(r, asteroids[i])
				break
			}
		}
	}
	fmt.Println(r)
}
