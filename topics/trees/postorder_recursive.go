package trees

func PostOrder(root *TreeNode) []int {
	items := []int{}

	var traversal func(*TreeNode, *[]int)

	traversal = func(root *TreeNode, items *[]int) {
		if root == nil {
			return
		}

		traversal(root.Left, items)
		traversal(root.Right, items)
		*items = append(*items, root.Val)

	}

	traversal(root, &items)
	return items
}
