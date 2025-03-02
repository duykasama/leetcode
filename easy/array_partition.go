package easy

import "fmt"

func ArrayPairSum(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		if i%2 != 0 {
			ans += nums[i]
		}
	}

	return ans
}
