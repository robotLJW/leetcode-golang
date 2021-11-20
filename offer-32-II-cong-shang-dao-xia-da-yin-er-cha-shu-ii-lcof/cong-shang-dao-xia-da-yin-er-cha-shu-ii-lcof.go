package offer_32_II_cong_shang_dao_xia_da_yin_er_cha_shu_ii_lcof

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root)
	count, nextCount := 1, 0
	array := []int{}
	for queue.Len() != 0 {
		tmpNode, _ := queue.Front().Value.(*TreeNode)
		array = append(array, tmpNode.Val)
		count--
		if tmpNode.Left != nil {
			queue.PushBack(tmpNode.Left)
			nextCount++
		}
		if tmpNode.Right != nil {
			queue.PushBack(tmpNode.Right)
			nextCount++
		}
		if count == 0 {
			ans = append(ans, array)
			count = nextCount
			nextCount = 0
			array = make([]int, 0)
		}
		queue.Remove(queue.Front())
	}
	return ans
}
