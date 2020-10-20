package _44_backspace_string_compare

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	t.Run("aaaa,aaaa", func(t *testing.T) {
		b := backspaceCompare("aaaa", "aaaa")
		assert.Equal(t, true, b)
	})

	t.Run("#aaaa#,aaaa", func(t *testing.T) {
		b := backspaceCompare("aaaa#", "aaaa")
		assert.Equal(t, false, b)
	})
}
