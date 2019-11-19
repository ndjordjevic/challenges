package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var t TreeNode
	createBSTFromSlice([]string{"5", "3", "6", "2", "4", "", "8", "1", "", "", "", "7", "9"}, &t)
	inOrder(&t)
}

func createBSTFromSlice(strings []string, root *TreeNode) {
	insertLevelOrder(strings, root, 0)
}

func insertLevelOrder(a []string, root *TreeNode, i int) {
	val, err := strconv.Atoi(a[i])
	if err != nil {
		return
	}

	if i < len(a) {
		t := &TreeNode{Val: val}
		root = t

		insertLevelOrder(a, root.Left, 2*i+1)
		insertLevelOrder(a, root.Right, 2*i+2)
	}
}

func inOrder(root *TreeNode) {
	if root != nil {
		inOrder(root.Left)
		fmt.Println(root.Val)
		inOrder(root.Right)
	}
}

//func increasingBST(root *TreeNode) *TreeNode {
//
//}
