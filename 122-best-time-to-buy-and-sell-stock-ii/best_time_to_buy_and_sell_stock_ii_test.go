package _22_best_time_to_buy_and_sell_stock_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	t.Run("7,1,5,3,6,4", func(t *testing.T) {
		data := []int{7, 1, 5, 3, 6, 4}
		expectValue := 7
		assert.Equal(t, expectValue, maxProfit(data))
	})
}
