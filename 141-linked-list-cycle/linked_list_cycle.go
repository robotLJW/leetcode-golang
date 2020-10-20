package _41_linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// hash
	if head == nil || head.Next == nil {
		return false
	}
	isVisited := make(map[*ListNode]bool)
	for head != nil {
		if ok, _ := isVisited[head]; ok {
			return true
		} else {
			isVisited[head] = true
			head = head.Next
		}
	}
	return false
}

func hasCycleTwo(head *ListNode) bool {
	// two point
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
