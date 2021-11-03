package _67_valid_perfect_square

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	t.Run("num is 1, should return true", func(t *testing.T) {
		ans := isPerfectSquare(1)
		if ans != true {
			t.Error("answer is wrong, should return true")
		}
	})

	t.Run("num is 16, should return true", func(t *testing.T) {
		ans := isPerfectSquare(16)
		if ans != true {
			t.Error("answer is wrong, should return true")
		}
	})

	t.Run("num is 14, should return false", func(t *testing.T) {
		ans := isPerfectSquare(14)
		if ans != false {
			t.Error("answer is wrong, should return false")
		}
	})

	t.Run("num is 4, should return true", func(t *testing.T) {
		ans := isPerfectSquare(4)
		if ans != true {
			t.Error("answer is wrong, should return true")
		}
	})
}
