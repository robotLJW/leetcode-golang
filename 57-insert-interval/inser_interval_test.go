package _7_insert_interval

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Insert(t *testing.T) {
	t.Run("", func(t *testing.T) {
		data := [][]int{
			{1, 3},
			{6, 9},
		}
		internal := []int{
			2, 5,
		}
		expect := [][]int{
			{1, 5},
			{6, 9},
		}
		assert.Equal(t, expect, insert(data, internal))
	})

	t.Run("", func(t *testing.T) {
		data := [][]int{
			{1, 2},
			{3, 5},
			{6, 7},
			{8, 10},
			{12, 16},
		}
		internal := []int{
			4, 8,
		}
		expect := [][]int{
			{1, 2},
			{3, 10},
			{12, 16},
		}
		assert.Equal(t, expect, insert(data, internal))
	})

	t.Run("", func(t *testing.T) {
		data := [][]int{
			{1, 5},
		}
		internal := []int{
			0, 3,
		}
		expect := [][]int{
			{0, 5},
		}
		assert.Equal(t, expect, insert(data, internal))
	})

	t.Run("", func(t *testing.T) {
		data := [][]int{
			{1, 5},
		}
		internal := []int{
			0, 0,
		}
		expect := [][]int{
			{0, 0},
			{1, 5},
		}
		assert.Equal(t, expect, insert(data, internal))
	})
}
