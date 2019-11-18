package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	value := root.Val
	return isEqual(root.Left, value) && isEqual(root.Right, value)
}

func isEqual(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val == val {
		return isEqual(root.Right, val) && isEqual(root.Left, val)
	}
	return false
}
