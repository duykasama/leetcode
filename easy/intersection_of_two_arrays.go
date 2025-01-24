package easy

func Intersection(nums1 []int, nums2 []int) []int {
	var result []int

	for _, num1 := range nums1 {
		contained := false
		for _, r := range result {
			if r == num1 {
				contained = true
				break
			}
		}

		if contained {
			continue
		}

		for _, num2 := range nums2 {
			if num2 == num1 {
				result = append(result, num1)
				break
			}
		}
	}

	return result
}
