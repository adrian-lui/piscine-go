package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	var state string
	if root == node {
		state = "equal"
	}
	if root.Data > node.Data {
		state = "left"
	}
	if root.Data < node.Data {
		state = "right"
	}
	if state == "equal" {
		newNode := TreeNode{}
		if root.Left != nil {
			newNode = *root.Left
			if root.Right != nil {
				newNode.Right = root.Right
			}
		}
		return &newNode
	}
	if state == "left" {
		if node.Left != nil {
			if node.Right != nil {
				node.Left.Right = node.Right
			} else {
				*node = *node.Left
			}
		}
	}
	if state == "right" {
		if node.Right != nil {
			if node.Left != nil {
				node.Right.Left = node.Left
			} else {
				*node = *node.Right
			}
		}
	}
	return root
}
