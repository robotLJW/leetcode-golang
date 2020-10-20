package _43_reorder_list

import "container/list"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Queue struct {
	list *list.List
}

func newQueue() *Queue {
	l := list.New()
	return &Queue{
		l,
	}
}

func (q *Queue) push(v interface{}) {
	q.list.PushBack(v)
}

func (q *Queue) popFront() interface{} {
	e := q.list.Front()
	if e != nil {
		q.list.Remove(e)
		return e.Value
	}
	return nil
}

func (q *Queue) popBack() interface{} {
	e := q.list.Back()
	if e != nil {
		q.list.Remove(e)
		return e.Value
	}
	return nil
}

func (q *Queue) isEmpty() bool {
	return q.list.Len() == 0
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	q := newQueue()
	for head != nil {
		q.push(head)
		head = head.Next
	}
	newHead := q.popFront().(*ListNode)
	currentNode := newHead
	for !q.isEmpty() {
		tempNode := q.popBack().(*ListNode)
		currentNode.Next = tempNode
		currentNode = tempNode
		if q.isEmpty() {
			break
		} else {
			tempNode = q.popFront().(*ListNode)
			currentNode.Next = tempNode
			currentNode = tempNode
		}
	}
	currentNode.Next = nil
	head = newHead
	return
}

func reorderListTwo(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}
