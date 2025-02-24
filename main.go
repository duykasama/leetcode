package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {
	a1 := []int{5, 5, 5}
	a2 := []int{2, 4, 2, 7}

	res := easy.MinimumBoxes(a1, a2)
	fmt.Println(res)
}
