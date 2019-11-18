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
	t := &TreeNode{Val: 4}
	fromArrayToTreeNode([]int{2, 7, 1, 3}, t)

	a := searchBST(t, 2)

	a.Traverse(func(r *TreeNode) {
		fmt.Println(r.Val)
	})
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	return nil
}

func fromArrayToTreeNode(a []int, root *TreeNode) {
	for _, v := range a {
		_ = root.Insert(v)
	}
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("%d", t.Val)
}

func (t *TreeNode) Traverse(f func(node *TreeNode)) {
	if t == nil {
		return
	}
	f(t)
	t.Left.Traverse(f)
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
