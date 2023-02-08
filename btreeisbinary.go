package piscine

var isBinary bool = true

func BTreeIsBinary(root *TreeNode) bool {
	if root.Left != nil {
		if root.Left.Data > root.Data {
			isBinary = false
		}
	}
	if root.Right != nil {
		if root.Right.Data < root.Data {
			isBinary = false
		}
	}
	if root.Left != nil {
		BTreeIsBinary(root.Left)
	}
	if root.Right != nil {
		BTreeIsBinary(root.Right)
	}
	return isBinary
}
