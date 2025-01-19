// https://leetcode.com/problems/maximum-odd-binary-number/description/
package easy

import "strings"

func MaximumOddBinaryNumber(s string) string {
	ones := strings.Count(s, "1")
	if ones == 0 {
		return s
	}

	return strings.Repeat("1", ones-1) + strings.Repeat("0", len(s)-ones) + "1"
}
