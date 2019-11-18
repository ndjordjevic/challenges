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

func main() {
	r1 := &TreeNode{Val: 1}
	fromArrayToTreeNode([]int{3, 2, 5}, r1)

	r2 := &TreeNode{Val: 2}
	fromArrayToTreeNode([]int{1, 3, 4, 7}, r2)

	r1.Traverse(func(r *TreeNode) {
		fmt.Println(r.Val)
	})

	fmt.Println()
	r2.Traverse(func(r *TreeNode) {
		fmt.Println(r.Val)
	})

	r := mergeTrees(r1, r2)

	r.Traverse(func(r *TreeNode) {
		fmt.Println(r.Val)
	})
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	return &TreeNode{t1.Val + t2.Val, mergeTrees(t1.Left, t2.Left), mergeTrees(t1.Right, t2.Right)}
}

func fromArrayToTreeNode(a []int, root *TreeNode) {
	for _, v := range a {
		_ = root.Insert(v)
	}
}
