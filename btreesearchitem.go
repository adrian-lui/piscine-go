package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	node := &TreeNode{Data: elem}
	if root.Data == elem {
		return root
	} else if root.Data > node.Data && root.Left != nil {
		node = BTreeSearchItem(root.Left, elem)
	} else if root.Data < node.Data && root.Right != nil {
		node = BTreeSearchItem(root.Right, elem)
	} else {
		return nil
	}
	return node
}
