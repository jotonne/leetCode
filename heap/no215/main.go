package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{2, 1}
	k := 1
	fmt.Println(findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {

	var h Intheap = Intheap(nums)
	heap.Init(&h)

	var x int
	count := len(nums) - k
	for h.Len() > 0 {
		if count < 0 {
			return x
		}
		x = heap.Pop(&h).(int)
		count--
	}
	return x
}

type Intheap []int

func (h Intheap) Len() int {
	return len(h)
}

func (h Intheap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Intheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Intheap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Intheap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}
