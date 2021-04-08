package _5_jump_game_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_jump(t *testing.T) {
	t.Run("[2 3 1 1 4] 2", func(t *testing.T) {
		data := []int{2, 3, 1, 1, 4}
		actual := jump(data)
		assert.Equal(t, 2, actual)
	})

	t.Run("[2 3 1 1 4] 2", func(t *testing.T) {
		data := []int{2, 3, 1, 1, 4}
		actual := jump2(data)
		assert.Equal(t, 2, actual)
	})
}
