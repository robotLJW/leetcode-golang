package _95_teemo_attacking

import "testing"

func TestFindPoisonedDuration(t *testing.T) {
	t.Run("timeSeries = [1,4], duration = 2 ,should return 4", func(t *testing.T) {
		timeSeries := []int{1, 4}
		ans := findPoisonedDuration(timeSeries, 2)
		if ans != 4 {
			t.Error("wrong answer , correct answer is 4")
		}
	})

	t.Run("timeSeries = [1,2], duration = 2 ,should return 3", func(t *testing.T) {
		timeSeries := []int{1, 2}
		ans:=findPoisonedDuration(timeSeries,2)
		if ans!=3{
			t.Error("wrong answer , correct answer is 3")
		}
	})

}
