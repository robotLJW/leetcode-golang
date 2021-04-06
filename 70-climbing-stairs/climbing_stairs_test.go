package _0_climbing_stairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ClimbStairs(t *testing.T) {
	t.Run("f(3)=3", func(t *testing.T) {
		ans := climbStairs(3)
		assert.Equal(t, 3, ans)
	})
}
