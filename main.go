package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {
	// input := []int{3, 6, 9}
	// result := easy.MinimumOperations(input)
	input := "1210"
	result := easy.DigitCount(input)
	fmt.Println("Result:", result)
}
