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
	//var t TreeNode
	//a := createBSTFromSlice([]string{"1", "2", "3", "4", "5", "6"}, &t)
	////printLevelOrder(a)

	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  6,
			Left: nil,
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	r := increasingBST(root)
	printInOrder(r)
}

func increasingBST(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	prev := dummy

	inorder(root, prev)

	prev.Left = nil
	prev.Right = nil
	//ret:=dummy.Right

	return prev.Right
}

func inorder(curr *TreeNode, prev *TreeNode) {
	if curr == nil {
		return
	}

	inorder(curr.Left, prev)
	prev.Left = nil
	prev.Right = curr
	prev = curr
	inorder(curr.Right, prev)
}

func printInOrder(root *TreeNode) {
	if root == nil {
		return
	}

	printInOrder(root.Left)
	fmt.Println(root.Val)
	printInOrder(root.Right)
}

func createBSTFromSlice(strings []string, root *TreeNode) *TreeNode {
	return insertLevelOrder(strings, root, 0)
}

func insertLevelOrder(arr []string, root *TreeNode, i int) *TreeNode {
	if i < len(arr) {
		val, err := strconv.Atoi(arr[i])

		if err != nil {
			return nil
		}

		t := &TreeNode{Val: val}
		root = t

		root.Left = insertLevelOrder(arr, root.Left, 2*i+1)
		root.Right = insertLevelOrder(arr, root.Right, 2*i+2)
	}

	return root
}

func printLevelOrder(root *TreeNode) {
	h := height(root)

	for i := 0; i < h; i++ {
		printGivenLevel(root, i)
	}

}

func printGivenLevel(root *TreeNode, i int) {
	if root == nil {
		return
	}
	if i == 0 {
		fmt.Println(root.Val)
	} else if i > 0 {
		printGivenLevel(root.Left, i-1)
		printGivenLevel(root.Right, i-1)
	}
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		lHeight := height(root.Left)
		rHeight := height(root.Right)

		if lHeight > rHeight {
			return lHeight + 1
		} else {
			return rHeight + 1
		}
	}
}
