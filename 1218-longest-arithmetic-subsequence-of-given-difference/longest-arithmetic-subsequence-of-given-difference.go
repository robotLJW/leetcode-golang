package _218_longest_arithmetic_subsequence_of_given_difference

// 超时
func longestSubsequence(arr []int, difference int) int {
	ans := 1
	for i := 0; i < len(arr); i++ {
		index := i
		length := 1
		for j := i + 1; j < len(arr); j++ {
			if arr[j]-arr[index] == difference {
				index = j
				length++
			}
		}
		if length > ans {
			ans = length
		}
	}
	return ans
}

func longestSubsequenceTwo(arr []int, difference int) int {
	dp := map[int]int{}
	ans := 1
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if dp[v] > ans {
			ans = dp[v]
		}
	}
	return ans
}
