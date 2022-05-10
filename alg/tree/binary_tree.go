package tree

type TreeNode struct {
	Val int

	Left  *TreeNode
	Right *TreeNode
}

func (td *TreeNode) InorderTraversal() []int {
	list := make([]int, 0)

	preOrder(td, &list)

	return list
}

func inorderTraversal(root *TreeNode) []int {
	list := make([]int, 0)

	preOrder(root, &list)

	return list
}

func preOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}

	*list = append(*list, root.Val)

	preOrder(root.Left, list)

	preOrder(root.Right, list)
}
