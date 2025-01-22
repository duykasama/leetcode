package main

import (
	"fmt"
	"leetcode/medium"
)

func main() {
	// input := []int{3, 6, 9}
	// result := easy.MinimumOperations(input)
	input := 2
	result := medium.NthPersonGetsNthSeat(input)
	fmt.Println("Result:", result)
}
