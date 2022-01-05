package _06_reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var currentNode *ListNode
	for head != nil {
		tmpNode := head
		head = head.Next
		tmpNode.Next = currentNode
		currentNode = tmpNode
	}
	return currentNode
}
