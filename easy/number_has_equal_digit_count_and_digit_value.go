package easy

func DigitCount(num string) bool {
	for i, n := range num {
		count := 0
		for _, r := range num {
			if i == int(r-'0') {
				count++
			}
		}

		if count != int(n-'0') {
			return false
		}
	}

	return true
}
