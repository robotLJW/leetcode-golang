package _30_minimum_absolute_difference_in_bst

import (
	"math"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	ans := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			ans = append(ans, node.Val)
			if node.Left != nil {
				dfs(node.Left)
			}
			if node.Right != nil {
				dfs(node.Right)
			}
		}
	}
	dfs(root)
	sort.Ints(ans)
	min := ans[len(ans)-1]
	for i := len(ans) - 1; i > 0; i-- {
		if min > ans[i]-ans[i-1] {
			min = ans[i] - ans[i-1]
		}
	}
	return min
}

// 中序遍历
func getMinimumDifferenceTwo(root *TreeNode) int {
	ans, pre := math.MaxInt64, -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
