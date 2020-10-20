package _16_populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 按层遍历
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		temp := queue
		queue = nil
		for i := 0; i < len(temp); i++ {
			if i+1 < len(temp) {
				temp[i].Next = temp[i+1]
			}
			if temp[i].Left != nil {
				queue = append(queue, temp[i].Left)
			}
			if temp[i].Right != nil {
				queue = append(queue, temp[i].Right)
			}
		}
	}
	return root
}
