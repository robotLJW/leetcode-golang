package _68_missing_number

import "testing"

func TestMissingNumber(t *testing.T) {
	t.Run("nums = [3,0,1],should return 2", func(t *testing.T) {
		data := []int{3, 0, 1}
		ans := missingNumber(data)
		if ans != 2 {
			t.Error("wrong answer, answer is 2")
		}
	})

	t.Run("nums = [0,1], should return 2", func(t *testing.T) {
		data := []int{0, 1}
		ans := missingNumber(data)
		if ans != 2 {
			t.Error("wrong answer, answer is 2")
		}
	})

	t.Run("nums = [9,6,4,2,3,5,7,0,1], should return 8", func(t *testing.T) {
		data := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
		ans := missingNumber(data)
		if ans != 8 {
			t.Error("wrong answer, answer is 8")
		}
	})

	t.Run("nums = [0], should return 1", func(t *testing.T) {
		data := []int{0}
		ans := missingNumber(data)
		if ans != 1 {
			t.Error("wrong answer, answer is 1")
		}
	})
}
