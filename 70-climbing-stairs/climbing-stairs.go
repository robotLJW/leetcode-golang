package _0_climbing_stairs

func climbStairs(n int) int {
	data := make(map[int]int)
	data[1] = 1
	data[2] = 2
	for i := 3; i <= n; i++ {
		data[i] = data[i-1] + data[i-2]
	}
	return data[n]
}
