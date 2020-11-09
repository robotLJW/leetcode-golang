package _73_k_closest_points_to_origin

import (
	"sort"
)

func kClosest(points [][]int, K int) [][]int {
	ans := make([][]int, K)
	sort.Slice(points, func(i, j int) bool {
		return points[i][0]*points[i][0]+points[i][1]*points[i][1] < points[j][0]*points[j][0]+points[j][1]*points[j][1]
	})
	for i := 0; i < K; i++ {
		ans[i] = make([]int, 2)
		ans[i][0] = points[i][0]
		ans[i][1] = points[i][1]
	}
	return ans
}
