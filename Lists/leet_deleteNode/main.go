package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node3 := ListNode{
		Val:  3,
		Next: nil,
	}

	node2 := ListNode{
		Val:  2,
		Next: &node3,
	}

	head := ListNode{
		Val:  1,
		Next: &node2,
	}

	fmt.Println("Before delete")
	printList(&head)
	deleteNode2(&node2)
	fmt.Println("After delete")
	printList(&head)
	fmt.Println()
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func deleteNode2(node *ListNode) {
	if node == nil {
		return
	}
	*node = *node.Next
}
