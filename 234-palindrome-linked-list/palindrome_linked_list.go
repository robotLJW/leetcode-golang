package _34_palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	data := make([]int, 0)
	for head != nil {
		val := head.Val
		data = append(data, val)
		head = head.Next
	}
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		if data[i] != data[j] {
			return false
		}
	}
	return true
}
