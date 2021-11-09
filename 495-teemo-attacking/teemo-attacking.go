package _95_teemo_attacking

func findPoisonedDuration(timeSeries []int, duration int) int {
	ans := duration
	for i := 1; i < len(timeSeries); i++ {
		ans += min(timeSeries[i]-timeSeries[i-1], duration)
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
