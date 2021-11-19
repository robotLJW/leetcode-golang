package _94_longest_harmonious_subsequence

func findLHS(nums []int) int {
	maps := make(map[int]int)
	ans := 0
	for i := 0; i < len(nums); i++ {
		if _, ok := maps[nums[i]]; ok {
			maps[nums[i]]++
		} else {
			maps[nums[i]] = 1
		}
	}
	for key, value := range maps {
		if v, ok := maps[key+1]; ok {
			if v+value > ans {
				ans = v + value
			}
		}
	}
	return ans
}
