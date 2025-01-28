package easy

func TwoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, num := range nums {
		dict[num] = i
	}

	for i, num := range nums {
		toFind := target - num
		idx, ok := dict[toFind]
		if ok && idx != i {
			return []int{i, idx}
		}
	}

	return []int{}
}
