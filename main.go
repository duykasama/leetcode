package main

import (
	"fmt"
	"leetcode/medium"
)

func main() {
	input := []int{1, 2, 3, 4}
	obj := medium.Constructor()

	for _, num := range input {
		obj.Add(num)
	}

	product := obj.GetProduct(3)
	fmt.Println(product)
}
