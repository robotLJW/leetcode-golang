package _7_combinations

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	var find func(path []int, start int)
	find = func(path []int, start int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			find(path, i+1)
			path = path[:len(path)-1]
		}
	}
	find([]int{}, 1)
	return ans
}
