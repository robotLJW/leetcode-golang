package _49_intersection_of_two_arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_intersection(t *testing.T) {
	t.Run("", func(t *testing.T) {
		data1 := []int{1, 2, 2, 1}
		data2 := []int{2, 2, 3, 4, 5}
		expectAns := []int{2}
		realAns := intersection(data1, data2)
		assert.Equal(t, expectAns, realAns)
	})
}
