package _77_squares_of_a_sorted_array

import "sort"

func sortedSquares(A []int) []int {
	for i := 0; i < len(A); i++ {
		A[i] = A[i] * A[i]
	}
	sort.Ints(A)
	return A
}

func sortedSquaresTwo(A []int) []int {
	ans := make([]int, len(A))
	i, j := 0, len(A)-1
	for index := len(ans) - 1; index >= 0; index-- {
		if v, w := A[i]*A[i], A[j]*A[j]; v < w {
			ans[index] = w
			j--
		} else {
			ans[index] = v
			i++
		}
	}
	return ans
}
