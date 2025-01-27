package easy

import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sArr := strings.Split(s, "")
	tArr := strings.Split(t, "")

	sort.Strings(sArr)
	sort.Strings(tArr)
	for i := 0; i < len(sArr); i++ {
		if sArr[i] != tArr[i] {
			return false
		}
	}
	return true
}
