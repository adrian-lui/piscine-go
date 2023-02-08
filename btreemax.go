package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root.Right == nil {
		return root
	}
	for root.Right != nil {
		root = root.Right
	}
	return root
}
