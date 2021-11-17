package _63_binary_tree_tilt

import "testing"

func TestFindTilt(t *testing.T) {
	t.Run("root = [1,2,3], should return 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		if findTilt(root) != 1 {
			t.Error("wrong answer, should return 1")
		}
	})
}
