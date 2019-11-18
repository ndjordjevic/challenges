package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	r := &TreeNode{Val: 10}
	fromArrayToTreeNode([]int{5, 15, 3, 7, 18}, r)

	r.Traverse(func(r *TreeNode) {
		fmt.Println(r.Val)
	})

	fmt.Println(rangeSumBST(r, 7, 15))
}

func fromArrayToTreeNode(a []int, root *TreeNode) {
	for _, v := range a {
		_ = root.Insert(v)
	}
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	var sum int

	switch {
	case root.Val < L:
		sum = rangeSumBST(root.Right, L, R)
	case root.Val > R:
		sum = rangeSumBST(root.Left, L, R)
	default:
		sum += root.Val
		sum += rangeSumBST(root.Left, L, R)
		sum += rangeSumBST(root.Right, L, R)
	}

	return sum
}

func (t *TreeNode) Traverse(f func(node *TreeNode)) {
	if t == nil {
		return
	}
	t.Left.Traverse(f)
	f(t)
	t.Right.Traverse(f)
}

func (n *TreeNode) Insert(value int) error {
	if n == nil {
		return errors.New("cannot insert a value into a nil tree")
	}

	switch {
	case value == n.Val:
		return nil
	case value < n.Val:
		if n.Left == nil {
			n.Left = &TreeNode{Val: value}
			return nil
		}
		return n.Left.Insert(value)
	case value > n.Val:
		if n.Right == nil {
			n.Right = &TreeNode{Val: value}
			return nil
		}
		return n.Right.Insert(value)
	}
	return nil
}
