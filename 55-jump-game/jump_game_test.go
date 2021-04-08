package _5_jump_game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canJump(t *testing.T) {
	t.Run("[3,2,1,0,0] false", func(t *testing.T) {
		data := []int{3, 2, 1, 0, 0}
		actual := canJump(data)
		assert.Equal(t, false, actual)
	})

	t.Run("[1,3,2,1,4] true", func(t *testing.T) {
		data := []int{1, 3, 2, 1, 4}
		actual := canJump(data)
		assert.Equal(t, true, actual)
	})
}
