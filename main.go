package main

import (
	"fmt"
	"leetcode/medium"
)

func main() {
	//input := []int{3, 2, 4}
	//input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	input := []string{"", "b", ""}
	result := medium.GroupAnagrams(input)
	fmt.Println(result)
}
