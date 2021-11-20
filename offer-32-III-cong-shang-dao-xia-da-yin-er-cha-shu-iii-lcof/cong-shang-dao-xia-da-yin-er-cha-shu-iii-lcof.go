package offer_32_III_cong_shang_dao_xia_da_yin_er_cha_shu_iii_lcof

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{
		list: list.New(),
	}
}

func (s *Stack) Push(node *TreeNode) {
	if s.list == nil {
		return
	}
	s.list.PushBack(node)
}

func (s *Stack) Pop() *TreeNode {
	if !s.Empty() {
		node := s.list.Back().Value.(*TreeNode)
		s.list.Remove(s.list.Back())
		return node
	}
	return nil
}

func (s *Stack) Empty() bool {
	return s.list.Len() == 0
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	stackOne := NewStack()
	stackTwo := NewStack()
	current, levelData := 0, make([]int, 0)
	stackOne.Push(root)
	for !stackOne.Empty() || !stackTwo.Empty() {
		if current == 0 {
			tmpNode := stackOne.Pop()
			levelData = append(levelData, tmpNode.Val)
			if tmpNode.Left != nil {
				stackTwo.Push(tmpNode.Left)
			}
			if tmpNode.Right != nil {
				stackTwo.Push(tmpNode.Right)
			}
			if stackOne.Empty() {
				ans = append(ans, levelData)
				levelData = make([]int, 0)
				current = 1 - current
			}
		} else {
			tmpNode := stackTwo.Pop()
			levelData = append(levelData, tmpNode.Val)
			if tmpNode.Right != nil {
				stackOne.Push(tmpNode.Right)
			}
			if tmpNode.Left != nil {
				stackOne.Push(tmpNode.Left)
			}
			if stackTwo.Empty() {
				ans = append(ans, levelData)
				levelData = make([]int, 0)
				current = 1 - current
			}
		}
	}
	return ans
}
