package _44_binary_tree_preorder_traversal

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack struct {
	list *list.List
}

func newStack() stack {
	l := list.New()
	return stack{
		l,
	}
}

func (s *stack) isEmpty() bool {
	return s.list.Len() == 0
}

func (s *stack) push(v interface{}) {
	s.list.PushBack(v)
}

func (s *stack) pop() interface{} {
	v := s.list.Back()
	if v != nil {
		s.list.Remove(v)
		return v.Value
	}
	return nil
}

// 非递归
func preorderTreversalTwo(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	s := newStack()
	s.push(root)
	for !s.isEmpty() {
		tempNode := s.pop().(*TreeNode)
		ans = append(ans, tempNode.Val)
		if tempNode.Right != nil {
			s.push(tempNode.Right)
		}
		if tempNode.Left != nil {
			s.push(tempNode.Left)
		}
	}
	return ans
}

// 递归
func preorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		} else {
			ans = append(ans, node.Val)
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
	}
	dfs(root)
	return ans
}
