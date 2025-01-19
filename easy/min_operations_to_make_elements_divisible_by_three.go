// https://leetcode.com/problems/find-minimum-operations-to-make-all-elements-divisible-by-three/description/
package easy

func MinimumOperations(nums []int) int {
	operations := 0
	for _, num := range nums {
		if num%3 == 0 {
			continue
		}
		operations++
	}

	return operations
}
