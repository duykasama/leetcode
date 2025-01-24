package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	result := easy.Intersection(nums1, nums2)
	for _, num := range result {
		fmt.Println(num)
	}
}
