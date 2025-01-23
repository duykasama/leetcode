package easy

func MinimumOperations_makeArrayAllZero(nums []int) int {
	ops := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}

	for _, num := range nums {
		if num == 0 {
			continue
		}
		ops++

		for i := range nums {
			if nums[i] == 0 {
				continue
			}
			nums[i] -= num
		}
	}

	return ops
}
