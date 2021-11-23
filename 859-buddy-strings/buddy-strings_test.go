package _59_buddy_strings

import "testing"

func TestBuddyStrings(t *testing.T) {
	t.Run("s = \"ab\", goal = \"ba\", should return true", func(t *testing.T) {
		if !buddyStrings("ab", "ba") {
			t.Error("wrong answer")
		}
	})

	t.Run("s = ab, goal = ab , should return false", func(t *testing.T) {
		if buddyStrings("ab", "ab") {
			t.Error("wrong answer")
		}
	})

}
