package _41_valid_mountain_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ValidMountainArray(t *testing.T) {
	t.Run("2,1", func(t *testing.T) {
		data := []int{2, 1}
		assert.Equal(t, false, validMountainArray(data))
	})

	t.Run("3,5,5", func(t *testing.T) {
		data := []int{3, 5, 5}
		assert.Equal(t, false, validMountainArray(data))
	})

	t.Run("0,3,2,1", func(t *testing.T) {
		data := []int{0, 3, 2, 1}
		assert.Equal(t, true, validMountainArray(data))
	})

}
