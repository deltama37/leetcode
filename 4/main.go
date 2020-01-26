package main

import "fmt"

func main() {
	a := []int{1, 3}
	b := []int{2}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(findMedianSortedArrays(a, b))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1i := 0
	nums2i := 0
	if (len(nums1)+len(nums2))%2 == 1 {
		for i := 0; i < (len(nums1)+len(nums2))/2; i++ {
			if nums1i == len(nums1)-1 {
				nums2i++
			}
			if nums2i == len(nums2)-1 {
				nums1i++
			}
			if nums1[nums1i] < nums2[nums2i] {
				nums1i++
			}
			if nums2[nums2i] < nums1[nums1i] {
				nums2i++
			}
		}
		if nums1[nums1i] < nums2[nums2i] {
			return float64(nums1[nums1i])
		} else {
			return float64(nums2[nums2i])
		}
	}
	return 0
}
