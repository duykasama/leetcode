package medium

func GroupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{{""}}
	}

	var output [][]string

	var innerArr []string
	for i, str := range strs {
		if str == "*" {
			continue
		}
		innerArr = append(innerArr, str)
		for j := i + 1; j < len(strs); j++ {
			if strs[j] == "*" {
				continue
			}
			anagrams := isAnagrams(innerArr[0], strs[j])
			if anagrams {
				innerArr = append(innerArr, strs[j])
				strs[j] = "*"
			}
		}
		output = append(output, innerArr)
		innerArr = []string{}
	}

	return output
}

func isAnagrams(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	arr := make([]int, 26)
	for _, r := range s {
		arr[r-'a'] += 1
	}
	for _, r := range t {
		arr[r-'a'] -= 1
	}
	for _, val := range arr {
		if val != 0 {
			return false
		}
	}
	return true
}
