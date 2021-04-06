package _35_lowest_common_ancestor_of_a_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pathP := getPath(root, p)
	pathQ := getPath(root, q)
	var ans *TreeNode
	for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i]; i++ {
		ans = pathP[i]
	}
	return ans
}

func getPath(root, targetNode *TreeNode) []*TreeNode {
	tmpNode := root
	path := make([]*TreeNode, 0)
	for tmpNode != targetNode {
		path = append(path, tmpNode)
		if tmpNode.Val > targetNode.Val {
			tmpNode = tmpNode.Left
		} else if tmpNode.Val < targetNode.Val {
			tmpNode = tmpNode.Right
		}
	}
	path = append(path, tmpNode)
	return path
}
