package _98_range_addition_ii

func maxCount(m int, n int, ops [][]int) int {
	minLine, minRow := m, n
	for i := 0; i < len(ops); i++ {
		if minLine > ops[i][0] {
			minLine = ops[i][0]
		}
		if minRow > ops[i][1] {
			minRow = ops[i][1]
		}
	}
	return minLine * minRow
}
