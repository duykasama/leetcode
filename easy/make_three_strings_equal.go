package easy

func FindMinimumOperations(s1 string, s2 string, s3 string) int {
	count := 0

	for i := 0; i < len(s1) && i < len(s2) && i < len(s3); i++ {
		if s1[i] == s2[i] && s1[i] == s3[i] {
			count++
			continue
		}

		break
	}

	if count == 0 {
		return -1
	}

	return len(s1) - count + len(s2) - count + len(s3) - count
}
