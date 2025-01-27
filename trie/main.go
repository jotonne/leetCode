package main

import (
	"fmt"
)

func main() {
	r := 'a'
	fmt.Println(r)
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, r := range word {
		index := r - 'a'
		if curr.children[index] == nil {
			curr.children[index] = &Trie{}
		}
		curr = curr.children[index]
	}
	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, r := range word {
		index := r - 'a'
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}
	return curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, r := range prefix {
		index := r - 'a'
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}
	return true
}
