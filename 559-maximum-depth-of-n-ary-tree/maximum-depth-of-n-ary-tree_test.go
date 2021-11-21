package _59_maximum_depth_of_n_ary_tree

import "testing"

func TestMaxDep(t *testing.T) {
	t.Run("root = [1,null,3,2,4], should return 2", func(t *testing.T) {
		root := &Node{
			Val: 1,
			Children: []*Node{
				&Node{
					Val: 3,
				},
				&Node{
					Val: 2,
				},
				&Node{
					Val: 4,
				},
			},
		}
		result := maxDepth(root)
		expect := 2

		if result != expect {
			t.Error("wrong answer, right answer is 2")
		}
	})
}
