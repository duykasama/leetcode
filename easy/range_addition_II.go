package easy

func MaxCount(m int, n int, ops [][]int) int {
	minX := m
	minY := n
	for _, op := range ops {
		if minX > op[0] {
			minX = op[0]
		}
		if minY > op[1] {
			minY = op[1]
		}
	}

	return minX * minY
}
