package _218_longest_arithmetic_subsequence_of_given_difference

import "testing"

func TestLongestSubsequence(t *testing.T) {
	t.Run("arr = [1,2,3,4], difference = 1, should return 4", func(t *testing.T) {
		data := []int{1, 2, 3, 4}
		ans := longestSubsequence(data, 1)
		if ans != 4 {
			t.Error("answer is wrong, answer is 4")
		}
	})

	t.Run("arr = [1,3,5,7], difference = 1, should return 1", func(t *testing.T) {
		data := []int{1, 3, 5, 7}
		ans := longestSubsequence(data, 1)
		if ans != 1 {
			t.Error("answer is wrong, answer is 1")
		}
	})

	t.Run("arr = [1,5,7,8,5,3,4,2,1], difference = -2, should return 4", func(t *testing.T) {
		data := []int{1, 5, 7, 8, 5, 3, 4, 2, 1}
		ans := longestSubsequence(data, -2)
		if ans != 4 {
			t.Error("answer is wrong, answer is 4")
		}
	})

}
