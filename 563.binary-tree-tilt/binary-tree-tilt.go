package _63_binary_tree_tilt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := dfs(node.Left)
		rightSum := dfs(node.Right)
		ans += abs(leftSum - rightSum)
		return leftSum+rightSum+node.Val
	}
	dfs(root)
	return
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
