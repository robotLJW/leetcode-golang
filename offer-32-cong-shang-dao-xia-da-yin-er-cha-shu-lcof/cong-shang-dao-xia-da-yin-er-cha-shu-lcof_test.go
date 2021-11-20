package offer_32_cong_shang_dao_xia_da_yin_er_cha_shu_lcof

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	t.Run("tree: [3,9,20,null,null,15,7] ,should return [3,9,20,15,7]", func(t *testing.T) {
		root := TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		}
		ans := []int{3, 9, 20, 15, 7}
		result := levelOrder(&root)
		if !reflect.DeepEqual(ans, result) {
			t.Error("wrong answer!")
		}
	})
}
