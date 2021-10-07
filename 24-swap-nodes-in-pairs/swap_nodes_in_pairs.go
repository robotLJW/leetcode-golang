package _4_swap_nodes_in_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tempHead := &ListNode{}
	tempHead.Next = head
	var swap func(pre, end, leftNode, RightNode *ListNode)
	swap = func(pre, end, leftNode, RightNode *ListNode) {
		pre.Next = RightNode
		RightNode.Next = leftNode
		leftNode.Next = end
	}
	temp := tempHead
	for temp.Next != nil && temp.Next.Next != nil {
		swap(temp, temp.Next.Next.Next, temp.Next, temp.Next.Next)
		temp = temp.Next.Next
	}
	return tempHead.Next
}
