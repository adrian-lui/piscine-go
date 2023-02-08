package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	iterator := root
	for iterator != nil {
		node := TreeNode{Data: data, Parent: iterator}
		if iterator.Data > data {
			if iterator.Left != nil {
				iterator = iterator.Left
			} else {
				iterator.Left = &node
				break
			}
		} else if iterator.Data < data {
			if iterator.Right != nil {
				iterator = iterator.Right
			} else {
				iterator.Right = &node
				break
			}
		}
	}
	return root
}
