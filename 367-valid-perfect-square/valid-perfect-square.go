package _67_valid_perfect_square

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	for start, end := 1, num/2; start <= end; {
		middle := (start + end) / 2
		ans := middle * middle
		if num == ans {
			return true
		} else if num > ans {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return false
}
