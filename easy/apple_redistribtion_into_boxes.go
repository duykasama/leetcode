package easy

func MinimumBoxes(apple []int, capacity []int) int {
	for i := 0; i < len(capacity); i++ {
		for j := i + 1; j < len(capacity); j++ {
			if capacity[i] < capacity[j] {
				temp := capacity[i]
				capacity[i] = capacity[j]
				capacity[j] = temp
			}
		}
	}

	total := 0
	for _, n := range apple {
		total += n
	}

	ans := 0
	for _, n := range capacity {
		ans++
		total -= n
		if total <= 0 {
			break
		}
	}
	return ans
}
