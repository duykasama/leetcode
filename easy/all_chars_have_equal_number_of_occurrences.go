package easy

func areOccurrencesEqual(s string) bool {
	appearances := make(map[rune]int)
	for _, c := range s {
		_, ok := appearances[c]
		if ok {
			appearances[c]++
		} else {
			appearances[c] = 1
		}
	}

	first := -1
	for _, val := range appearances {
		if first == -1 {
			first = val
			continue
		}

		if first != val {
			return false
		}
	}

	return true
}
