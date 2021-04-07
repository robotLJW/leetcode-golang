package _21_best_time_to_buy_and_sell_stock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProfit(t *testing.T) {

	t.Run("test [7 1 5 3 6 4]", func(t *testing.T) {
		data := []int{7, 1, 5, 3, 6, 4}
		ans := maxProfit(data)
		expectValue := 5
		assert.Equal(t, expectValue, ans)
	})

	t.Run("test [7,6,4,3,1]", func(t *testing.T) {
		data := []int{7, 6, 4, 3, 1}
		ans := maxProfit(data)
		expectValue := 0
		assert.Equal(t, expectValue, ans)
	})
}
