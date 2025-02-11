package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {
	s1 := "abc"
	s2 := "abb"
	s3 := "ab"
	result := easy.FindMinimumOperations(s1, s2, s3)
	fmt.Println(result)
}
