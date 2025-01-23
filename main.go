package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {
	// input := []int{3, 6, 9}
	// result := easy.MinimumOperations(input)
	input := []int{2, 3, 1}
	result := easy.SubarraySum(input)
	fmt.Println("Result:", result)
}
