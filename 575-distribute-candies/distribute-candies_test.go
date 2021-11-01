package _75_distribute_candies

import (
	"errors"
	"testing"
)

func TestDistributeCandies(t *testing.T) {
	t.Run("data is [1,1,2,2,3,3], should return 3", func(t *testing.T) {
		data := []int{1, 1, 2, 2, 3, 3}
		ans := distributeCandies(data)
		if ans != 3 {
			t.Error(errors.New("answer is 3"))
		}
	})

	t.Run("data is [1,1,2,3], should return 2", func(t *testing.T) {
		data := []int{1, 1, 2, 3}
		ans := distributeCandies(data)
		if ans != 2 {
			t.Error(errors.New("answer is 2"))
		}
	})

	t.Run("data is [1,1,1,1], should return 1", func(t *testing.T) {
		data := []int{1, 1, 1, 1}
		ans:=distributeCandies(data)
		if ans!=1{
			t.Error(errors.New("answer is 1"))
		}
	})
}
