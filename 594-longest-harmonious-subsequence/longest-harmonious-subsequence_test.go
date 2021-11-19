package _94_longest_harmonious_subsequence

import "testing"

func TestFindLHS(t *testing.T){
	t.Run("data [1,1,1,1] should return 0", func(t *testing.T) {
		data:=[]int{1,1,1,1}
		if findLHS(data)!=0{
			t.Error("wrong answer, should return 0")
		}
	})

	t.Run("data [1,2,3,4] should return 2", func(t *testing.T) {
		data:=[]int{1,2,3,4}
		if findLHS(data)!=2{
			t.Error("wrong answer, should return 2")
		}
	})
}
