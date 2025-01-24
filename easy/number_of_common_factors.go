package easy

func CommonFactors(a int, b int) int {
	if a > b {
		temp := a
		a = b
		b = temp
	}

	var result []int
	for i := 1; i <= a; i++ {
		if a%i == 0 && b%i == 0 {
			result = append(result, i)
		}
	}

	return len(result)
}
