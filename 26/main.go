package main

import (
	"fmt"
)

func main() {
	a := []int{1, 1, 2}
	res := removeDuplicates(a)
	fmt.Println(a[:res])
}

func removeDuplicates(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[count] != nums[i] {
			count++
			nums[count] = nums[i]
		}
	}
	return count + 1
}
