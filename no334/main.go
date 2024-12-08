package main

func main() {
	nums := []int{20, 100, 10, 12, 5, 13}
	if len(nums) <= 2 {
		// return false
	}

	bak := 2147483647
	bak2 := 2147483647
	for i := 0; i < len(nums); i++ {
		if nums[i] < bak {
			bak = nums[i]
			continue
		}
		if nums[i] < bak2 {
			bak2 = nums[i]
			continue
		}
		// return true
	}
	// return false
}
