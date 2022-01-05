package _06_reverse_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseList(t *testing.T) {
	listNode := ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	list := reverseList(&listNode)
	assert.Equal(t, 2, list.Val)
	assert.Equal(t, 1, list.Next.Val)

}
