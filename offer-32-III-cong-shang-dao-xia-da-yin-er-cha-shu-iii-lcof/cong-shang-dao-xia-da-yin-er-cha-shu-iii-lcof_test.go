package offer_32_III_cong_shang_dao_xia_da_yin_er_cha_shu_iii_lcof

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	t.Run("data: [3,9,20,null,null,15,7] ,should return [\n  [3],\n  [20,9],\n  [15,7]\n]", func(t *testing.T) {
		root := &TreeNode{
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
		result := levelOrder(root)
		expect := [][]int{
			{3},
			{20, 9},
			{15, 7},
		}
		if !reflect.DeepEqual(result, expect) {
			t.Error("wrong answer.")
		}
	})
}
