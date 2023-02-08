package piscine

func findLevel(root *TreeNode, f func(...interface{}) (int, error), h int) {
	if root == nil {
		return
	}
	if h == 1 {
		f(root.Data)
	}
	if h > 1 {
		findLevel(root.Left, f, h-1)
		findLevel(root.Right, f, h-1)
	}
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	height := BTreeLevelCount(root)

	for i := 1; i <= height; i++ {
		findLevel(root, f, i)
	}
}
