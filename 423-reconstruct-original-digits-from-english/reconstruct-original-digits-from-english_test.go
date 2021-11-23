package _23_reconstruct_original_digits_from_english

import (
	"reflect"
	"testing"
)

func TestOriginalDigits(t *testing.T) {
	t.Run("s = \"owoztneoer\", should return 012", func(t *testing.T) {
		result := originalDigits("owoztneoer")
		expect := "012"
		if !reflect.DeepEqual(result, expect) {
			t.Error("wrong answer")
		}
	})

}
