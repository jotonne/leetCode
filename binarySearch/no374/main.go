package main

import "fmt"

const pick = 6

func main() {
	fmt.Println(guessNumber(7))
}

func guessNumber(n int) int {
	low := 0
	high := n

	for low <= high {
		guessed := (low + high) / 2
		resultGuess := guess(guessed)
		switch resultGuess {
		case 0:
			return guessed
		case -1:
			high = guessed - 1
		case 1:
			low = guessed + 1
		}
	}

	return -1
}

func guess(num int) int {
	if num == pick {
		return 0
	}
	if num > pick {
		return -1
	} else {
		return 1
	}
}
