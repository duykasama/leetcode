package easy

func SubarraySum(nums []int) int {
	sum := 0
	for i := range nums {
		start := i - nums[i]
		if start < 0 {
			start = 0
		}
		for j := start; j <= i; j++ {
			sum += nums[j]
		}
	}

	return sum
}
