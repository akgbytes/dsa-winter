package main

import (
	"fmt"
)

func main() {

	nums := []int{-2, 2, 4, 4, 4, 4, 5, 5}
	ans := 4
	fmt.Println("fn returned: ", removeDuplicates(nums), " Correct ans: ", ans)
}

func removeDuplicates(nums []int) int {
	pos := 1
	i := 1
	for i < len(nums) {
		if nums[i] != nums[i-1] {
			nums[pos] = nums[i]
			pos++
		}
		i++
	}
	return pos
}
