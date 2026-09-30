package trees

func InOrder(root *TreeNode) []int {
	items := []int{}

	var traversal func(*TreeNode, *[]int)
	traversal = func(root *TreeNode, items *[]int) {
		if root == nil {
			return
		}

		traversal(root.Left, items)
		*items = append(*items, root.Val)
		traversal(root.Right, items)
	}

	traversal(root, &items)
	return items
}
