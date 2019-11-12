package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{}
	createList(head, 1, 2, 3, 4)
	printList(head)
	printList(head)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head)
		head = head.Next
	}

}

func createList(head *ListNode, nums ...int) {
	var prev *ListNode

	for _, val := range nums {
		if prev == nil {
			head.Val = val
			prev = head
			continue
		}

		newNode := &ListNode{val, nil}
		prev.Next = newNode
		prev = newNode
	}
}
