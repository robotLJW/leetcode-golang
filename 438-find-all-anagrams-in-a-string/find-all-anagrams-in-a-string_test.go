package _38_find_all_anagrams_in_a_string

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	t.Run("s = \"cbaebabacd\", p = \"abc\" ,should return [0,6]", func(t *testing.T) {
		result := findAnagrams("cbaebabacd", "abc")
		expect := []int{0, 6}
		if !reflect.DeepEqual(result, expect) {
			t.Error("wrong answer, should return [0,6]")
		}
	})

	t.Run("s = \"abab\", p = \"ab\" ,should return [0,1,2]", func(t *testing.T) {
		result := findAnagrams("abab", "ab")
		expect := []int{0, 1, 2}
		if !reflect.DeepEqual(result, expect) {
			t.Error("wrong answer, should return [0,1,2]")
		}
	})
}
