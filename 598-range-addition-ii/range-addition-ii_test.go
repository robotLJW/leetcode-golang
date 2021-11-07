package _98_range_addition_ii

import "testing"

func TestMaxCount(t *testing.T) {
	t.Run("data m = 3,n = 3 ,operations = [[2,2],[3,3]], should return 4", func(t *testing.T) {
		opts := [][]int{{2, 2}, {3, 3}}
		ans := maxCount(3, 3, opts)
		if ans != 4 {
			t.Error("wrong answer, answer is 4")
		}
	})
}
