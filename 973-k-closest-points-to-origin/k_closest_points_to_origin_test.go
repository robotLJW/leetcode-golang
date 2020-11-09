package _73_k_closest_points_to_origin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_kClosest(t *testing.T) {
	t.Run("", func(t *testing.T) {
		data := [][]int{
			{1, 3},
			{-2, 2},
		}
		expect := [][]int{
			{-2, 2},
		}
		assert.Equal(t, expect, kClosest(data, 1))
	})
}
