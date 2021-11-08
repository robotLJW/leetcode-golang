package _99_bulls_and_cows

import "testing"

func TestGetHint(t *testing.T) {
	t.Run("secret: 1807, guess: 7810, should return 1A3B", func(t *testing.T) {
		secret, guess := "1807", "7810"
		ans := "1A3B"
		if getHint(secret, guess) != ans {
			t.Error("wrong answer should return 1A3B")
		}
	})

	t.Run("secret: 11, guess: 10, should return 1A1B", func(t *testing.T) {
		secret, guess := "11", "10"
		ans:="1A0B"
		if getHint(secret,guess)!=ans{
			t.Error("wrong answer should return 1A0B")
		}
	})
}
