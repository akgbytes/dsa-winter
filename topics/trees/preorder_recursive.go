package trees

func PreOrder(root *TreeNode) []int {
	items := []int{}

	var traversal func(*TreeNode, *[]int)

	traversal = func(root *TreeNode, items *[]int) {
		if root == nil {
			return
		}

		*items = append(*items, root.Val)
		traversal(root.Left, items)
		traversal(root.Right, items)
	}

	traversal(root, &items)
	return items
}
