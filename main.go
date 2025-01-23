package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {
	// input := []int{3, 6, 9}
	// result := easy.MinimumOperations(input)
	input := []int{0}
	result := easy.MinimumOperations_makeArrayAllZero(input)
	fmt.Println("Result:", result)
}
