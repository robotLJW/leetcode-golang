package _356_sort_integers_by_the_number_of_1_bits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SortByBits(t *testing.T) {
	t.Run("0,1,2,3,4,5,6,7,8", func(t *testing.T) {
		data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
		ans := []int{0, 1, 2, 4, 8, 3, 5, 6, 7}
		assert.Equal(t, ans, sortByBits(data))
	})
}
