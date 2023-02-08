package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	rplc.Parent = node.Parent
	*node = *rplc
	return root
}
