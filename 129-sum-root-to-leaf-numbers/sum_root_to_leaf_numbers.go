package _29_sum_root_to_leaf_numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var ans int
	var dfs func(node *TreeNode, preValue int)
	dfs = func(node *TreeNode, preValue int) {
		if node == nil {
			return
		}
		preValue = preValue*10 + node.Val
		if node.Left == nil && node.Right == nil {
			ans = ans + preValue
		}
		dfs(node.Left, preValue)
		dfs(node.Right, preValue)
	}
	dfs(root, 0)
	return ans
}
