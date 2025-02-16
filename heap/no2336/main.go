package main

import (
	"container/heap"
	"fmt"
)

func main() {
	s := Constructor()
	s.AddBack(2)
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
}

type SmallestInfiniteSet struct {
	intheap  *intheap
	existVal map[int]bool
}

func Constructor() SmallestInfiniteSet {
	nums := make([]int, 1000)
	existVal := make(map[int]bool)
	for i := range nums {
		nums[i] = i + 1
		existVal[i] = true
	}

	var h intheap = intheap(nums)
	heap.Init(&h)
	return SmallestInfiniteSet{
		intheap:  &h,
		existVal: existVal,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.intheap.Len() <= 0 {
		return 0
	}

	min := heap.Pop(this.intheap).(int)
	this.existVal[min] = false
	return min
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if exist := this.existVal[num]; exist {
		return
	}
	heap.Push(this.intheap, num)
	this.existVal[num] = true
}

type intheap []int

func (h intheap) Len() int {
	return len(h)
}

func (h intheap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intheap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *intheap) Pop() any {
	result := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return result
}
