package piscine

func leftD(root *TreeNode) int {
	if root.Left != nil {
		return 1 + BTreeLevelCount(root.Left)
	}
	if root.Right != nil {
		return 1 + BTreeLevelCount(root.Right)
	}
	return 1
}

func rightD(root *TreeNode) int {
	if root.Left != nil {
		return 1 + BTreeLevelCount(root.Left)
	}
	if root.Right != nil {
		return 1 + BTreeLevelCount(root.Right)
	}
	return 1
}

func BTreeLevelCount(root *TreeNode) int {
	left := 0
	right := 0
	if root.Left != nil {
		left = leftD(root.Left)
	}
	if root.Right != nil {
		right = rightD(root.Right)
	}
	if left > right {
		return 1 + left
	} else {
		return 1 + right
	}
}
