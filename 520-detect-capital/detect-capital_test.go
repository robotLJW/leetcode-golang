package _20_detect_capital

import "testing"

func TestDetectCapitalUse(t *testing.T) {
	t.Run("word is USA, answer should return true", func(t *testing.T) {
		word := "USA"
		if !detectCapitalUse(word) {
			t.Error("wrong answer, should return true")
		}
	})

	t.Run("word is Google, answer should return true", func(t *testing.T) {
		word := "Google"
		if !detectCapitalUse(word){
			t.Error("wrong answer, should return true")
		}
	})

	t.Run("word is FlaG, answer should return false", func(t *testing.T) {
		word:="FlaG"
		if detectCapitalUse(word){
			t.Error("wrong answer, should return false")
		}
	})
}
