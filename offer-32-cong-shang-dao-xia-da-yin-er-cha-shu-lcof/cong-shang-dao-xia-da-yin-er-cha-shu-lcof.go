package offer_32_cong_shang_dao_xia_da_yin_er_cha_shu_lcof

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	ans := []int{}
	queue := list.New()
	if root == nil {
		return ans
	}
	queue.PushBack(root)
	for queue.Len() != 0 {
		topNode, _ := queue.Front().Value.(*TreeNode)
		if topNode.Left != nil {
			queue.PushBack(topNode.Left)
		}
		if topNode.Right != nil {
			queue.PushBack(topNode.Right)
		}
		ans = append(ans, topNode.Val)
		queue.Remove(queue.Front())
	}
	return ans
}
