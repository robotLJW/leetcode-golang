package _356_sort_integers_by_the_number_of_1_bits

import (
	"sort"
)

func sortByBits(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		x, y := a[i], a[j]
		cx, cy := countOneNumb(x), countOneNumb(y)
		return cx < cy || cx == cy && x < y
	})
	return a
}


func countOneNumb(numb int) int {
	var c int
	for ; numb > 0; numb /= 2 {
		c += numb % 2
	}
	return c
}
