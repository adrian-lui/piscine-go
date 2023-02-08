package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}
