package _59_maximum_depth_of_n_ary_tree

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{
		list: list.New(),
	}
}

func (q *Queue) Push(node *Node) {
	q.list.PushBack(node)
}

func (q *Queue) Pop() *Node {
	if !q.Empty() {
		node := q.list.Front().Value.(*Node)
		q.list.Remove(q.list.Front())
		return node
	}
	return nil
}

func (q *Queue) Empty() bool {
	return q.list.Len() == 0
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	queueOne := NewQueue()
	queueTwo := NewQueue()
	deep, current := 0, 0
	queueOne.Push(root)
	for !queueOne.Empty() || !queueTwo.Empty() {
		if current == 0 {
			tmpNode := queueOne.Pop()
			childNodes := tmpNode.Children
			for _, node := range childNodes {
				queueTwo.Push(node)
			}
			if queueOne.Empty() {
				deep++
				current = 1 - current
			}
		} else {
			tmpNode := queueTwo.Pop()
			childNodes := tmpNode.Children
			for _, node := range childNodes {
				queueOne.Push(node)
			}
			if queueTwo.Empty() {
				deep++
				current = 1 - current
			}
		}
	}
	return deep
}
