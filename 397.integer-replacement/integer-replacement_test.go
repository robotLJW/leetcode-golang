package _97_integer_replacement

import "testing"

func TestIntegerReplacement(t *testing.T){
	t.Run("n=7 ,should return 4", func(t *testing.T) {
		if integerReplacement(7)!=4{
			t.Error("wrong answer, should return 4")
		}
	})
}
